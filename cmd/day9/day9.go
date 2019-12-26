package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"github.com/enjean/advent-of-code-2019/internal/intcode"
)

func ExecuteProgramWithInput(program []intcode.IPType, inputs []intcode.IPType) []intcode.IPType {
	computer := intcode.CreateCompleteComputer("")

	go func() { computer.Run(program) }()

	for _, input := range inputs {
		computer.Input <- input
	}

	var outputs []intcode.IPType
	for output := range computer.Output {
		outputs = append(outputs, output)
	}
	return outputs
}

func main() {
	program := intcode.ParseProgram(adventutil.Parse(9)[0])
	part1 := ExecuteProgramWithInput(program, []intcode.IPType{1})
	fmt.Printf("Part 1 %v\n", part1)

	part2 := ExecuteProgramWithInput(program, []intcode.IPType{2})
	fmt.Printf("Part 2 %v\n", part2)
}
