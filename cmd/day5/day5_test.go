package main

import "testing"

func TestExecuteProgramWithInput(t *testing.T) {
	program := []int{3, 0, 4, 0, 99}
	input := 987

	result := ExecuteProgramWithInput(program, input)

	if result[0] != input {
		t.Errorf("Program did not produce the same output as input")
	}
}
