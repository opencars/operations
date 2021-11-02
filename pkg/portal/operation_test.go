package portal_test

import (
	"testing"

	"github.com/opencars/operations/pkg/portal"
	"github.com/stretchr/testify/assert"
)

func TestFixBrand(t *testing.T) {
	flagTests := []struct {
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

			outBrand, outModel := portal.FixBrandModel(test.inBrand, test.inModel)
			assert.Equal(t, test.outBrand, outBrand)
			assert.Equal(t, test.outModel, outModel)
		})
	}
}
