package main

import (
	"reflect"
	"testing"
)

func TestWordFrequencyCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "basic word count",
			input: "hello world hello",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			name:  "case insensitive counting",
			input: "Go GO go",
			expected: map[string]int{
				"go": 3,
			},
		},

		{
			name:  "remove punctuation",
			input: "hello, world! hello.",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			name:     "empty string",
			input:    "",
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := WordFrequencyCount(test.input)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("WordFrequencyCount(%q) = %v; want %v",
					test.input,
					result,
					test.expected)
			}
		})
	}

}
