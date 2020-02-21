package model_test

import (
	"testing"

	"github.com/opencars/operations/pkg/model"
)

func TestFixBrand(t *testing.T) {
	var flagtests = []struct {
		inBrand, inModel   string
		outBrand, outModel string
	}{
		{"TESLA  MODEL X", "MODEL  X", "TESLA", "MODEL X"},
	}

	for _, tt := range flagtests {
		t.Run(tt.inBrand, func(t *testing.T) {
			outBrand, outModel := model.FixBrand(tt.inBrand, tt.inModel)
			if outBrand != tt.outBrand || outModel != tt.outModel {
				t.Errorf("got %q, want %q", outBrand, tt.outBrand)
				t.Errorf("got %q, want %q", outModel, tt.outModel)
			}
		})
	}

}
