package model

import (
	"testing"
	"time"
)

// TestResource return special resource for testing purposes.
func TestResource(t *testing.T) *Resource {
	t.Helper()

	return &Resource{
		UID:          "1235678-1234-1234-1234-000123456789",
		Name:         "example.csv",
		URL:          "https://data.gov.ua/dataset/00000000-0000-0000-0000-00000000000/resource/1235678-1234-1234-1234-000123456789/download/example.csv",
		LastModified: time.Now(),
	}
}

// TestOperation return special operations for testing purposes.
func TestOperation(t *testing.T) *Operation {
	t.Helper()

	fuel := "ЕЛЕКТРО"
	ownWeight := 2485.0
	totalWeight := 3021.0
	vin := "5YJXCCE40GF010543"

	return &Operation{
		Person:      "P",
		RegAddress:  nil,
		RegCode:     410,
		Reg:         "ПЕРЕРЕЄСТРАЦІЯ ПРИ ЗАМІНІ НОМЕРНОГО ЗНАКУ",
		Date:        "2019-06-01",
		DepCode:     12290,
		Dep:         "ТСЦ 8041",
		Brand:       "TESLA",
		Model:       "MODEL X",
		VIN:         &vin,
		Year:        2016,
		Color:       "ЧОРНИЙ",
		Kind:        "ЛЕГКОВИЙ",
		Body:        "УНІВЕРСАЛ-B",
		Purpose:     "ЗАГАЛЬНИЙ",
		Fuel:        &fuel,
		Capacity:    nil,
		OwnWeight:   &ownWeight,
		TotalWeight: &totalWeight,
		Number:      "АА9359РС",
	}
}
