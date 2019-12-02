package main

import "testing"

func TestRun(t *testing.T) {
	//1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).
	//2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
	//2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
	//1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.
	tests := []struct {
		program  []int
		expected []int
	}{
		{
			program:  []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			expected: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			program:  []int{1,0,0,0,99},
			expected: []int{2,0,0,0,99},
		},
		{
			program:  []int{2,3,0,3,99},
			expected: []int{2,3,0,6,99},
		},
		{
			program:  []int{2,4,4,5,99,0},
			expected: []int{2,4,4,5,99,9801},
		},
		{
			program:  []int{1,1,1,4,99,5,6,0,99},
			expected: []int{30,1,1,4,2,5,6,0,99},
		},
	}
	for _, test := range tests {
		Run(test.program)
		for i, pVal := range test.program {
			expectedVal := test.expected[i]
			if pVal != expectedVal {
				t.Errorf("Program %v not equal to expected %v", test.program, test.expected)
				break
			}
		}
	}
}
