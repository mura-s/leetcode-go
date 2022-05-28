package problem3

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	inputs := []struct {
		input  string
		output int
	}{
		{
			input:  "abcabcbb",
			output: 3,
		},
		{
			input:  "bbbbb",
			output: 1,
		},
		{
			input:  "pwwkew",
			output: 3,
		},
		{
			input:  "a",
			output: 1,
		},
	}

	for _, input := range inputs {
		result := lengthOfLongestSubstring(input.input)
		if !reflect.DeepEqual(result, input.output) {
			t.Errorf("got %v, want %v", result, input.output)
		}
	}
}
