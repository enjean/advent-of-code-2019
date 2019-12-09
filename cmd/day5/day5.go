package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"github.com/enjean/advent-of-code-2019/internal/intcode"
)

func ExecuteProgramWithInput(program []intcode.IPType, input int) []int {
	computer := intcode.CreateComputer("", map[int]intcode.Instruction{
		1: intcode.Add,
		2: intcode.Multiply,
		3: intcode.Save,
		4: intcode.PrintFunc,
		5: intcode.JumpIfTrue,
		6: intcode.JumpIfFalse,
		7: intcode.LessThan,
		8: intcode.Equals,
	})

	go func() { computer.Input <- intcode.IPType(input) }()
	go func() { computer.Run(program) }()

	var outputs []int
	for output := range computer.Output {
		outputs = append(outputs, int(output))
	}
	return outputs
}

func main() {
	program := intcode.ParseProgram(adventutil.Parse(5)[0])

	part1Outputs := ExecuteProgramWithInput(program, 1)
	fmt.Printf("Part 1: %v\n", part1Outputs)

	part2Outputs := ExecuteProgramWithInput(program, 5)
	fmt.Printf("Part 2: %v\n", part2Outputs)
}
