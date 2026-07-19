package main

import "testing"

func TestPalindromeCheck(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "simple palindrome",
			input:    "madam",
			expected: true,
		},
		{
			name:     "not a palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := palindromeCheck(test.input)

			if result != test.expected {
				t.Errorf(
					"palindromeCheck(%q) = %v; want %v",
					test.input,
					result,
					test.expected,
				)
			}
		})
	}
}
