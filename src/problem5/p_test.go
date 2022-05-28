package problem5

import (
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	inputs := []struct {
		input  string
		output string
	}{
		{
			input:  "babad",
			output: "bab",
		},
		{
			input:  "cbbd",
			output: "bb",
		},
	}

	for _, input := range inputs {
		result := longestPalindrome(input.input)
		if !reflect.DeepEqual(result, input.output) {
			t.Errorf("got %v, want %v", result, input.output)
		}
	}
}
