package worker

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/opencars/govdata"

	"github.com/opencars/operations/pkg/domain"
	"github.com/opencars/operations/pkg/domain/model"
	"github.com/opencars/operations/pkg/logger"
)

// Worker is responsible for processing incoming data.
type Worker struct {
	store  domain.Store
	parser domain.Parser
}

// New returns new instance of worker.
func New(s domain.Store, p domain.Parser) *Worker {
	return &Worker{
		store:  s,
		parser: p,
	}
}

// Process dispatches event as to a handler.
func (w *Worker) Process(ctx context.Context, packageID string) error {
	modified, err := w.modifiedResources(ctx)
	if err != nil {
		return err
	}

	resources := govdata.SubscribePackage(packageID, modified)

	for {
		select {
		case resource, ok := <-resources:
			if !ok {
				return nil
			}

			log := logger.WithFields(logger.Fields{
				"id":   resource.ID,
				"name": resource.Name,
			})

			log.WithFields(logger.Fields{
				"mime_type":     resource.MimeType,
				"revisions":     len(resource.Revisions),
				"package_id":    resource.PackageID,
				"last_modified": resource.LastModified,
				"url":           resource.URL,
			}).Infof("resource event received")

			if err := w.handle(ctx, &resource); err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (w *Worker) reader(ctx context.Context, event *govdata.Resource) (io.ReadCloser, error) {
	var r io.ReadCloser

	switch event.MimeType {
	case "application/zip":
		reader, err := w.unzip(ctx, event.URL)
		if err != nil {
			return nil, err
		}

		r = reader

		logger.Debugf("archive unzipped")
	case "text/csv":
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, event.URL, http.NoBody)
		if err != nil {
			return nil, err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		r = resp.Body
	default:
		return nil, errors.New("invalid mime type")
	}

	return r, nil
}

func (w *Worker) handle(ctx context.Context, event *govdata.Resource) error {
	current, err := w.store.Resource().FindByUID(ctx, event.ID)
	if err != nil && err != model.ErrNotFound {
		return nil
	}

	if !errors.Is(err, model.ErrNotFound) {
		affected, err := w.store.Operation().DeleteByResourceID(ctx, current.ID)
		if err != nil {
			return err
		}

		logger.WithFields(logger.Fields{
			"affected": affected,
		}).Infof("entities were successfully deleted")
	}

	resource := model.Resource{
		UID:          event.ID,
		Name:         event.Name,
		URL:          event.URL,
		LastModified: time.Unix(0, 0),
	}

	logger.Infof("resource modification time was reset")

	if err := w.store.Resource().Create(ctx, &resource); err != nil {
		return err
	}

	reader, err := w.reader(ctx, event)
	if err != nil {
		return err
	}

	start := time.Now()

	if err := w.parser.Parse(ctx, &resource, reader); err != nil {
		return err
	}

	logger.WithFields(logger.Fields{
		"duration": time.Since(start),
	}).Infof("finished parsing resource")

	resource.LastModified = event.LastModified.Time
	if err := w.store.Resource().Update(ctx, &resource); err != nil {
		return err
	}

	return nil
}

func (w *Worker) unzip(ctx context.Context, url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	reader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, err
	}

	// Read all the files from zip archive
	if len(reader.File) != 1 {
		return nil, fmt.Errorf("zip: invalid amount of files: %d", len(reader.File))
	}

	return reader.File[0].Open()
}

func (w *Worker) modifiedResources(ctx context.Context) (map[string]time.Time, error) {
	resources, err := w.store.Resource().All(ctx)
	if err != nil {
		return nil, err
	}

	modified := make(map[string]time.Time)
	for _, r := range resources {
		modified[r.UID] = r.LastModified
	}

	return modified, nil
}
