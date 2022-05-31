package problem13

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	inputs := []struct {
		input  string
		output int
	}{
		{
			input:  "III",
			output: 3,
		},
		{
			input:  "LVIII",
			output: 58,
		},
		{
			input:  "MCMXCIV",
			output: 1994,
		},
	}

	for _, input := range inputs {
		result := romanToInt(input.input)
		if !reflect.DeepEqual(result, input.output) {
			t.Errorf("got %v, want %v", result, input.output)
		}
	}
}
