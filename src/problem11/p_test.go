package problem11

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	inputs := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			output: 49,
		},
		{
			input:  []int{1, 1},
			output: 1,
		},
		{
			input:  []int{2, 3, 4, 5, 18, 17, 6},
			output: 17,
		},
	}

	for _, input := range inputs {
		result := maxArea(input.input)
		if !reflect.DeepEqual(result, input.output) {
			t.Errorf("got %v, want %v", result, input.output)
		}
	}
}
