package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"github.com/enjean/advent-of-code-2019/internal/intcode"
)

func ExecuteProgramWithInput(program []intcode.IPType, inputs []intcode.IPType) []intcode.IPType {
	computer := intcode.CreateComputer("", map[int]intcode.Instruction{
		1: intcode.Add,
		2: intcode.Multiply,
		3: intcode.Save,
		4: intcode.PrintFunc,
		5: intcode.JumpIfTrue,
		6: intcode.JumpIfFalse,
		7: intcode.LessThan,
		8: intcode.Equals,
		9: intcode.AdjustRelativeBase,
	})

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
}
