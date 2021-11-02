package csv_test

import (
	"strings"
	"testing"

	"github.com/opencars/operations/pkg/csv"
	"github.com/opencars/operations/pkg/domain/model"
	"github.com/stretchr/testify/assert"
)

type Example struct {
	Name    string  `csv:"name"`
	Surname string  `csv:"surname"`
	Age     int     `csv:"age"`
	Balance float64 `csv:"balance"`
}

type Example2 struct {
	Name    string `csv:"name"`
	Surname string `csv:"surname"`
}

func TestRowUnmarshaller_Unmarshal(t *testing.T) {
	fields := map[string]int{
		"name":    3,
		"surname": 2,
		"age":     1,
		"balance": 0,
	}

	actual := Example{}

	expected := Example{
		Name:    "john",
		Surname: "doe",
		Age:     30,
		Balance: 150.75,
	}

	d := csv.NewRowDecoder(fields)

	err := d.Decode([]string{
		"150.75", "30", "doe", "john",
	}, &actual)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestReader_ReadBulk(t *testing.T) {
	example := "name,surname\njohn1,doe1\njohn2,doe2"

	r := csv.NewReader(strings.NewReader(example), ',')

	arr := []Example{}

	err := r.ReadBulk(1, &arr)
	assert.NoError(t, err)
}

func TestReader_ReadBulk2(t *testing.T) {
	example := "\"name\";\"surname\"\n\"john1\";\"doe1\"\n\"john2\";\"doe2\""

	r := csv.NewReader(strings.NewReader(example), ';')

	arr := []model.Operation{}

	err := r.ReadBulk(1, &arr)
	assert.NoError(t, err)
}
