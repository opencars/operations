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
		fmt.Println(rowValue)

		switch fv.Kind() {
		case reflect.String:
			fv.SetString(rowValue)
		case reflect.Int, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(rowValue, 10, 64)
			if err != nil {
				return err
			}

			fv.SetInt(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(rowValue, 64)
			if err != nil {
				return err
			}

			fv.SetFloat(x)
		default:
			log.Println("Unsupported", fv.Type())
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

	xv = xv.Elem()
	xt = xt.Elem()

	if xv.Kind() != reflect.Slice {
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
			return err
		}

		xv = reflect.Append(xv, decoded.Elem())
	}

	return nil
}
