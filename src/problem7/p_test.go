package problem7

import (
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	inputs := []struct {
		input  int
		output int
	}{
		{
			input:  123,
			output: 321,
		},
		{
			input:  -123,
			output: -321,
		},
		{
			input:  120,
			output: 21,
		},
	}

	for _, input := range inputs {
		result := reverse(input.input)
		if !reflect.DeepEqual(result, input.output) {
			t.Errorf("got %v, want %v", result, input.output)
		}
	}
}
