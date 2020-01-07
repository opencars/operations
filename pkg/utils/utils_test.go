package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opencars/operations/pkg/utils"
)

func TestAtoi(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "1",
			output: 1,
		},
		{
			input:  "9223372036854775807",
			output: 9223372036854775807,
		},
		{
			input:  "-9223372036854775807",
			output: -9223372036854775807,
		},
	}

	for _, test := range tests {
		out, err := utils.Atoi(&test.input)
		assert.NoError(t, err)
		assert.Equal(t, test.output, *out)
	}
}

func TestAtof(t *testing.T) {
	tests := []struct {
		input  string
		output float64
	}{
		{
			input:  "1.0",
			output: 1,
		},
		{
			input:  "123456789.123456789",
			output: 123456789.123456789,
		},
		{
			input:  "-123456789.123456789",
			output: -123456789.123456789,
		},
	}

	for _, test := range tests {
		out, err := utils.Atof(&test.input)
		assert.NoError(t, err)
		assert.Equal(t, test.output, *out)
	}
}
