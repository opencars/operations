package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type RowDecoder struct {
	fields map[string]int
}

func NewRowDecoder(fields map[string]int) *RowDecoder {
	return &RowDecoder{
		fields: fields,
	}
}

func (u *RowDecoder) Decode(row []string, x interface{}) error {
	xv := reflect.ValueOf(x)
	xt := reflect.TypeOf(x)

	if xv.Kind() != reflect.Ptr {
		return fmt.Errorf("should be a pointer")
	}

	xv = xv.Elem()
	xt = xt.Elem()

	for i := 0; i < xv.NumField(); i++ {
		ft := xt.Field(i)
		fv := xv.Field(i)

		tag := strings.Split(ft.Tag.Get("csv"), ",")
		if len(tag) == 0 {
			continue
		}

		fieldIndex, ok := u.fields[tag[0]]
		if !ok {
			continue
		}

		rowValue := row[fieldIndex]
		rowValue = strings.Trim(rowValue, "\"")

		if rowValue == "" || rowValue == "NULL" {
			continue
		}

		fvc := fv
		var isPtr bool

		if fvc.Kind() == reflect.Ptr && fvc.IsNil() {
			fvc = reflect.New(fvc.Type().Elem()).Elem()
			isPtr = true
		}

		switch fvc.Kind() {
		case reflect.String:
			fvc.SetString(rowValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(rowValue, 10, 64)
			if err != nil {
				return err
			}

			fvc.SetInt(x)
		case reflect.Float32, reflect.Float64:
			rowValue = strings.ReplaceAll(rowValue, ",", ".")

			x, err := strconv.ParseFloat(rowValue, 64)
			if err != nil {
				return err
			}

			fvc.SetFloat(x)
		default:
			log.Println("Unsupported", fvc.Type())
		}

		if isPtr {
			fv.Set(fvc.Addr())
		}
	}

	return nil
}

type Reader struct {
	*csv.Reader

	decoder *RowDecoder
}

func NewReader(r io.Reader, delimeter rune) *Reader {
	csvr := csv.NewReader(r)
	csvr.Comma = delimeter

	return &Reader{
		Reader:  csvr,
		decoder: nil,
	}
}

func (r *Reader) ReadBulk(amount int, x interface{}) error {
	xv := reflect.ValueOf(x)
	xt := reflect.TypeOf(x)

	if xv.Kind() != reflect.Ptr {
		return fmt.Errorf("non-pointer %v", xv.Type())
	}

	xve := xv.Elem()
	xt = xt.Elem()

	if xve.Kind() != reflect.Slice {
		return fmt.Errorf("can't fill non-slice value")
	}

	for i := 0; i < amount; i++ {
		row, err := r.Read()
		if errors.Is(err, io.EOF) {
			return err
		}

		if err != nil {
			return err
		}

		if r.decoder == nil {
			fields := make(map[string]int)

			for j, f := range row {
				fields[f] = j
			}

			r.decoder = NewRowDecoder(fields)

			i--
			continue
		}

		decoded := reflect.New(xt.Elem())
		if err := r.decoder.Decode(row, decoded.Interface()); err != nil {
			log.Print(row, err)
		}

		xve = reflect.Append(xve, decoded.Elem())
	}

	return nil
}
