package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/intcode"
	"testing"
)

func TestExecuteProgramWithInput1(t *testing.T) {
	program := []intcode.IPType{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99}
	result := ExecuteProgramWithInput(program, nil)
	fmt.Printf("Should be itself %d", result)
}
