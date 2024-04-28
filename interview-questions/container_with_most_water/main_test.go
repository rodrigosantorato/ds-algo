package main

import (
	"testing"
)

/**
Find container with most water and return the area of said container
each number in the input array is a height, and the width is the diff from one index to the other

scenario 1:
biggest container is 9 to 7, the height of the container is 7 and the width is 5
i.e. (index of 7 = 5) - (index of 9 = 0) = 5
input: [9,1,2,3,4,7] res: 35

scenario 2:
biggest container is 30 to 50, the height of the container is 30 and the width is 3
input: [9,30,2,3,50,7] res: 90
**/

func TestFindAreaOfBiggestContainer(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "happy scenario",
			input:    []int{9, 1, 2, 3, 4, 7},
			expected: 35,
		},
		{
			name:     "another happy scenario",
			input:    []int{9, 30, 2, 3, 50, 7},
			expected: 90,
		},
		{
			name:     "empty scenario",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "single num scenario",
			input:    []int{9},
			expected: 0,
		},
	}

	for _, test := range tests {
		res := findAreaOfBiggestContainerBruteForce(test.input)
		if res != test.expected {
			t.Errorf("%v failed with. Expected %v but got %v instead", test.name, test.expected, res)
		}
	}
}
