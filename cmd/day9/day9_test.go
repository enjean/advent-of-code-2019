package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/intcode"
	"testing"
)

func TestExecuteProgramWithInput1(t *testing.T) {
	program := []intcode.IPType{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	result := ExecuteProgramWithInput(program, nil)
	fmt.Printf("Should be itself %d\n", result)
}

func TestExecuteProgramWithInput2(t *testing.T) {
	program := []intcode.IPType{1102,34915192,34915192,7,4,7,99,0}
	result := ExecuteProgramWithInput(program, nil)
	fmt.Printf("Should be a 16 digit number %d\n", result)
}

func TestExecuteProgramWithInput3(t *testing.T) {
	program := []intcode.IPType{104,1125899906842624,99}
	result := ExecuteProgramWithInput(program, nil)
	if result[0] != 1125899906842624 {
		t.Error("Failed")
	}
}
