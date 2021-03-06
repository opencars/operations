package bulkreader

import (
	"io"
)

type Reader interface {
	Read() (record []string, err error)
}

type BulkReader struct {
	reader Reader
}

func New(r Reader) *BulkReader {
	return &BulkReader{
		reader: r,
	}
}

func (r *BulkReader) ReadBulk(amount int) ([][]string, error) {
	result := make([][]string, 0, amount)

	for i := 0; i < amount; i++ {
		record, err := r.reader.Read()
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
