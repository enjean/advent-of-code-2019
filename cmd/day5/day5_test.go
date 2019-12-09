package main

import (
	"github.com/enjean/advent-of-code-2019/internal/intcode"
	"testing"
)

func TestExecuteProgramWithInput(t *testing.T) {
	program := []intcode.IPType{3, 0, 4, 0, 99}
	input := 987

	result := ExecuteProgramWithInput(program, input)

	if result[0] != input {
		t.Errorf("Program did not produce the same output as input")
	}
}
