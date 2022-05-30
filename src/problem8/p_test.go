package problem8

import (
	"reflect"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	inputs := []struct {
		input  string
		output int
	}{
		{
			input:  "42",
			output: 42,
		},
		{
			input:  "   -42",
			output: -42,
		},
		{
			input:  "4193 with words",
			output: 4193,
		},
		{
			input:  "",
			output: 0,
		},
		{
			input:  "+-12",
			output: 0,
		},
		{
			input:  "  0000000000012345678",
			output: 12345678,
		},
		{
			input:  "00000-42a1234",
			output: 0,
		},
		{
			input:  ".1",
			output: 0,
		},
	}

	for _, input := range inputs {
		result := myAtoi(input.input)
		if !reflect.DeepEqual(result, input.output) {
			t.Errorf("got %v, want %v", result, input.output)
		}
	}
}
