package main

import (
	"testing"
)

func TestCalculateThrusterSignal(t *testing.T) {
	tests := []struct {
		phases   [5]int
		program  []int
		expected int
	}{
		{
			[5]int{4, 3, 2, 1, 0},
			[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			43210,
		},
		{
			[5]int{0, 1, 2, 3, 4},
			[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			54321,
		},
		{
			[5]int{1, 0, 4, 3, 2},
			[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			65210,
		},
	}

	for _, test := range tests {
		//fmt.Printf("********* Test %d\n", test.expected)
		result := CalculateThrusterSignal(test.phases, test.program)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}

func TestOptimalThrusterSignal(t *testing.T) {
	tests := []struct {
		program  []int
		expectedPhases   [5]int
		expected int
	}{
		{
			[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			[5]int{4, 3, 2, 1, 0},
			43210,
		},
		{
			[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			[5]int{0, 1, 2, 3, 4},
			54321,
		},
		{
			[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			[5]int{1, 0, 4, 3, 2},
			65210,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			phases, thruster := OptimalThrusterSignal(test.program)
			if phases != test.expectedPhases && thruster != test.expected {
				t.Errorf("Expected %d for %v, got %v %d", test.expected, test.expectedPhases, phases, thruster)
			}
		})
	}
}
