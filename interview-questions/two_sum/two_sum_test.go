package main

import (
	"slices"
	"testing"
)

// Find the two numbers that sum to the target and return their indexes
// [1, 2, 3, 4, 5]  target = 9   result = [3,4]

type Test struct {
	name     string
	input    []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	tests := []Test{
		{
			name:     "happy ending",
			input:    []int{1, 2, 3, 4, 5},
			target:   9,
			expected: []int{3, 4},
		},
		{
			name:     "another happy ending",
			input:    []int{1, 2, 10, 3, 4, 5, 20},
			target:   30,
			expected: []int{2, 6},
		},
		{
			name:     "no answer",
			input:    []int{1, 2, 3, 4, 5},
			target:   30,
			expected: []int{},
		},
		{
			name:     "empty input",
			input:    []int{},
			target:   9,
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{10},
			target:   10,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := twoSum(tt.input, tt.target)

			if !slices.Equal(tt.expected, res) {
				t.Errorf("%v failed. expected was : %v but instead got %v\n", tt.name, tt.expected, res)
			}
		})
	}
}
