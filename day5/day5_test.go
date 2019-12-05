package main

import "testing"

func TestIsImmediateMode(t *testing.T) {
	tests := []struct{
		opcode int
		argNumber int
		expectedIsImmediateMode bool
	}{
		{1002, 1, false},
		{1002, 2, true},
		{1002, 3, false},
		{3, 1, false},
		{3, 2, false},
		{3, 3, false},
		{99, 1, false},
		{99, 2, false},
		{99, 3, false},
		{199, 1, true},
		{199, 2, false},
		{199, 3, false},
		{1199, 1, true},
		{1199, 2, true},
		{1199, 3, false},
		{10099, 1, false},
		{10099, 2, false},
		{10099, 3, true},
	}
	for _, test := range tests {
		result := isImmediateMode(test.opcode, test.argNumber)
		if result != test.expectedIsImmediateMode {
			t.Errorf("Op %d arg %d expected %v was %v", test.opcode, test.argNumber, test.expectedIsImmediateMode, result)
		}
	}
}
