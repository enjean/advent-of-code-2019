package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"github.com/enjean/advent-of-code-2019/internal/intcode"
	"strconv"
	"strings"
)

func ExecuteProgramWithInput(program []int, input int) []int {
	executable := make([]int, len(program))
	copy(executable, program)

	computer := intcode.CreateComputer(map[int]func(intcode.Computer, []int, int) int{
		1: intcode.Add,
		2: intcode.Multiply,
		3: intcode.Save,
		4: intcode.PrintFunc,
		5: intcode.JumpIfTrue,
		6: intcode.JumpIfFalse,
		7: intcode.LessThan,
		8: intcode.Equals,
	})

	go func() { computer.Input <- input }()
	go func() { computer.Run(executable) }()

	var outputs []int
	for output := range computer.Output {
		outputs = append(outputs, output)
	}
	return outputs
}

func main() {
	programString := adventutil.Parse(5)[0]
	partsStrings := strings.Split(programString, ",")
	var program []int
	for _, partString := range partsStrings {
		asInt, _ := strconv.Atoi(partString)
		program = append(program, asInt)
	}

	part1Outputs := ExecuteProgramWithInput(program, 1)
	fmt.Printf("Part 1: %v\n", part1Outputs)

	part2Outputs := ExecuteProgramWithInput(program, 5)
	fmt.Printf("Part 2: %v\n", part2Outputs)
}
