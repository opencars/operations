package worker

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/opencars/govdata"
	"github.com/opencars/operations/pkg/logger"
	"github.com/opencars/operations/pkg/model"
	"github.com/opencars/operations/pkg/store"
)

const (
	mappers   = 10
	reducers  = 10
	shufflers = 10 // Don't change. It closes channel.
	batchSize = 2500
)

type handlerCSV struct {
	reader *csv.Reader
}

func (h *handlerCSV) readN(amount int) ([][]string, error) {
	result := make([][]string, 0)

	for i := 0; i < amount; i++ {
		record, err := h.reader.Read()
		if err == io.EOF {
			return result, err
		}

		if err != nil {
			return nil, err
		}

		result = append(result, record)
	}

	return result, nil
}

func shuffler(wg *sync.WaitGroup, input chan model.Operation, output chan []model.Operation) {
	operations := make([]model.Operation, 0, batchSize)

	for {
		operation, ok := <-input
		if !ok {
			if len(operations) != 0 {
				output <- operations
			}
			wg.Done()
			break
		}

		if len(operations) < batchSize {
			operations = append(operations, operation)
			continue
		}

		output <- operations

		// TODO: Find another way to find too much memory allocation.
		operations = make([]model.Operation, 0, batchSize)
	}
}

func mapper(id int64, wg *sync.WaitGroup, input chan []string, output chan model.Operation) {
	for {
		msg, ok := <-input
		if !ok {
			wg.Done()
			break
		}

		oper, err := model.OperationFromGov(msg)
		if err != nil {
			log.Println(err)
			continue
		}
		oper.ResourceID = id

		output <- *oper
	}
}

func reducer(wg *sync.WaitGroup, store store.Store, input chan []model.Operation) {
	for {
		operations, ok := <-input
		if !ok {
			wg.Done()
			break
		}

		if err := store.Operation().Add(operations...); err != nil {
			log.Fatal(err)
		}

		log.Printf("Done: %d\n", len(operations))
	}
}

func mapperDispatcher(handler handlerCSV, output chan []string) {
	for {
		msgs, err := handler.readN(2500)

		if err == nil || err == io.EOF {
			for _, msg := range msgs {
				output <- msg
			}
		}

		if err == io.EOF {
			close(output)
			break
		}

		if err != nil {
			log.Println(err)
			close(output)
			break
		}
	}
}

// Worker is responsible for processing incoming data.
type Worker struct {
	store store.Store
}

// New returns new instance of worker.
func New(store store.Store) *Worker {
	return &Worker{
		store: store,
	}
}

// Process dispatches event as to a handler.
func (w *Worker) Process(event interface{}) error {
	switch v := event.(type) {
	case govdata.Resource:
		if err := w.handleResouce(&v); err != nil {
			return err
		}
	case govdata.Revision:
		if err := w.handleRevision(&v); err != nil {
			return err
		}
	default:
		return errors.New("invalid type")
	}

	return nil
}

func (w *Worker) handleResouce(resource *govdata.Resource) error {
	logger.Info("Started parsing resource %s", resource.Name)

	res := model.Resource{
		UID:          resource.ID,
		Name:         resource.Name,
		LastModified: resource.LastModified.Time.UTC(),
		URL:          resource.URL,
	}

	err := w.store.Resource().Create(&res)

	if err != nil {
		return err
	}

	reader, err := w.unzip(resource.URL)
	if err != nil {
		return err
	}
	defer reader.Close()

	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'

	// Skip header line.
	if _, err := csvReader.Read(); err != nil {
		return err
	}

	start := time.Now()
	handler := handlerCSV{reader: csvReader}
	rows := make(chan []string, 100000)
	operations := make(chan model.Operation, 100000)
	batches := make(chan []model.Operation, 10000)

	mapperWg := sync.WaitGroup{}
	shufflersWg := sync.WaitGroup{}
	reducersWg := sync.WaitGroup{}

	for i := 0; i < reducers; i++ {
		reducersWg.Add(1)
		go reducer(&reducersWg, w.store, batches)
	}

	for i := 0; i < shufflers; i++ {
		shufflersWg.Add(1)
		go shuffler(&shufflersWg, operations, batches)
	}

	for i := 0; i < mappers; i++ {
		mapperWg.Add(1)
		go mapper(res.ID, &mapperWg, rows, operations)
	}

	go mapperDispatcher(handler, rows)

	mapperWg.Wait()

	// Close channel.
	time.Sleep(time.Second)
	close(operations)

	shufflersWg.Wait()

	// Close channel.
	time.Sleep(time.Second)
	close(batches)

	time.Sleep(time.Second)
	// Wait for reducers.
	reducersWg.Wait()

	logger.Info("Execution time: %s", time.Since(start))
	logger.Info("Finished parsing resource %s", resource.Name)

	return nil
}

func (w *Worker) handleRevision(revision *govdata.Revision) error {
	logger.Info("Started parsing revisoin %s", revision.Name)

	resource, err := govdata.ResourceShow(revision.ResourceID)
	if err != nil {
		return err
	}

	// Find resource.
	res, err := w.store.Resource().FindByUID(revision.ResourceID)
	if err != nil {
		return err
	}

	// Delete all operations.
	if err := w.store.Operation().DeleteByResourceID(res.ID); err != nil {
		return err
	}

	// Add all operations again.
	logger.Info("Started parsing resource %s", resource.Name)
	reader, err := w.unzip(revision.URL)
	if err != nil {
		return err
	}
	defer reader.Close()

	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'

	// Skip header line.
	if _, err := csvReader.Read(); err != nil {
		return err
	}

	start := time.Now()
	handler := handlerCSV{reader: csvReader}
	rows := make(chan []string, 100000)
	operations := make(chan model.Operation, 100000)
	batches := make(chan []model.Operation, 10000)

	mapperWg := sync.WaitGroup{}
	shufflersWg := sync.WaitGroup{}
	reducersWg := sync.WaitGroup{}

	for i := 0; i < reducers; i++ {
		reducersWg.Add(1)
		go reducer(&reducersWg, w.store, batches)
	}

	for i := 0; i < shufflers; i++ {
		shufflersWg.Add(1)
		go shuffler(&shufflersWg, operations, batches)
	}

	for i := 0; i < mappers; i++ {
		mapperWg.Add(1)
		go mapper(res.ID, &mapperWg, rows, operations)
	}

	go mapperDispatcher(handler, rows)

	mapperWg.Wait()

	// Close channel.
	time.Sleep(time.Second)
	close(operations)

	shufflersWg.Wait()

	// Close channel.
	time.Sleep(time.Second)
	close(batches)

	time.Sleep(time.Second)
	// Wait for reducers.
	reducersWg.Wait()

	logger.Info("Execution time: %s", time.Since(start))
	logger.Info("Finished parsing resource %s", resource.Name)

	// Update resource last_modified.
	res.URL = revision.URL
	res.LastModified = resource.LastModified.Time
	if err := w.store.Resource().Update(res); err != nil {
		return err
	}

	return nil
}

func (w *Worker) unzip(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	reader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, err
	}

	// Read all the files from zip archive
	if len(reader.File) != 1 {
		return nil, fmt.Errorf("invalid amount of files in zip: %d", len(reader.File))
	}

	return reader.File[0].Open()
}

func (w *Worker) ModifiedResources() map[string]time.Time {
	resources, err := w.store.Resource().All()
	if err != nil {
		logger.Fatal(err)
	}

	modified := make(map[string]time.Time)
	for _, r := range resources {
		modified[r.UID] = r.LastModified
	}

	return modified
}
