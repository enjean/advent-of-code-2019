package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"strconv"
	"strings"
)

func initialize(program []int, noun, verb int) []int{
	newProgram := make([]int, len(program))
	copy(newProgram, program)
	newProgram[1] = noun
	newProgram[2] = verb
	return newProgram
}

func Run(program []int) {
	ip := 0
	for {
		opcode := program[ip]
		if opcode == 99 {
			break
		}
		input1Position := program[ip + 1]
		input1 := program[input1Position]
		input2Position := program[ip + 2]
		input2 := program[input2Position]
		outputPosition := program[ip + 3]
		if opcode == 1 {
			program[outputPosition] = input1 + input2
		}
		if opcode == 2 {
			program[outputPosition] = input1 * input2
		}
		ip += 4
	}
}

func main() {
	programString := adventutil.Parse(2)[0]
	partsStrings := strings.Split(programString, ",")
	var program []int
	for _, partString := range partsStrings {
		asInt, _ := strconv.Atoi(partString)
		program = append(program, asInt)
	}

	// replace position 1 with the value 12 and replace position 2 with the value 2
	part1Program := initialize(program, 12, 2)
	Run(part1Program)
	fmt.Printf("Part 1: %d\n", part1Program[0])

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			nvProgram := initialize(program, noun, verb)
			Run(nvProgram)
			if nvProgram[0] == 19690720 {
				fmt.Printf("Part 2: %d\n", 100 * noun + verb)
				break
			}
		}
	}
}
