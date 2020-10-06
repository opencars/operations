package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opencars/operations/pkg/model"
)

func TestFixBrand(t *testing.T) {
	var flagTests = []struct {
		inBrand, inModel   string
		outBrand, outModel string
	}{
		{
			"TESLA  MODEL X",
			"MODEL  X", "TESLA",
			"MODEL X",
		},
	}

	for i := range flagTests {
		test := flagTests[i]
		t.Run(test.inBrand, func(t *testing.T) {
			outBrand, outModel := model.FixBrandModel(test.inBrand, test.inModel)
			assert.Equal(t, test.outBrand, outBrand)
			assert.Equal(t, test.outModel, outModel)
		})
	}
}
