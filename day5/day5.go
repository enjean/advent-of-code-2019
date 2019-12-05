package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/adventutil"
	"strconv"
	"strings"
)

type Computer struct {
	ops map[int]func([]int, int) int
}

func add(program []int, ip int) int {
	//fmt.Println(program[ip:ip+4])
	arg1 := getValue(program, ip, 1)
	arg2 := getValue(program, ip, 2)
	dest := program[ip+3]
	//fmt.Printf("Setting %d to %d + %d\n", dest, arg1, arg2)
	program[dest] = arg1 + arg2
	return ip + 4
}

func multiply(program []int, ip int) int {
	arg1 := getValue(program, ip, 1)
	arg2 := getValue(program, ip, 2)
	dest := program[ip+3]
	program[dest] = arg1 * arg2
	return ip + 4
}

func save(program []int, ip int) int {
	//fmt.Println(program[ip:ip+2])
	dest := program[ip+1]
	var input int
	fmt.Println("Enter input:")
	_, err := fmt.Scan(&input)
	if err != nil {
		panic("Error reading input")
	}
	program[dest] = input
	return ip + 2
}

func printFunc(program []int, ip int) int {
	fmt.Println(program[ip:ip+2])
	value := getValue(program, ip, 1)
	fmt.Println(value)
	return ip + 2
}

func jumpIfTrue(program []int, ip int) int {
	arg1 := getValue(program, ip, 1)
	if arg1 != 0 {
		return getValue(program, ip, 2)
	}
	return ip + 3
}

func jumpIfFalse(program []int, ip int) int {
	arg1 := getValue(program, ip, 1)
	if arg1 == 0 {
		return getValue(program, ip, 2)
	}
	return ip + 3
}

func lessThan(program []int, ip int) int {
	arg1 := getValue(program, ip, 1)
	arg2 := getValue(program, ip, 2)
	dest := program[ip+3]
	result := 0
	if arg1 < arg2 {
		result = 1
	}
	program[dest] = result
	return ip + 4
}

func equals(program []int, ip int) int {
	arg1 := getValue(program, ip, 1)
	arg2 := getValue(program, ip, 2)
	dest := program[ip+3]
	result := 0
	if arg1 == arg2 {
		result = 1
	}
	program[dest] = result
	return ip + 4
}

func getValue(program []int, ip, argNum int) int {
	opcode := program[ip]
	argValue := program[ip+argNum]
	if isImmediateMode(opcode, argNum) {
		return argValue
	}
	return program[argValue]
}

func isImmediateMode(opcode, argNum int) bool {
	return (opcode/(adventutil.Pow10(argNum+1)))%10 == 1
}

func (c Computer) Run(program []int) {
	ip := 0
	for {
		opcode := program[ip] % 100
		fmt.Printf("IP %d op %d\n", ip, program[ip])
		if opcode == 99 {
			break
		}
		ip = c.ops[opcode](program, ip)
	}
}

func main() {
	computer := Computer{
		map[int]func([]int, int) int{
			1: add,
			2: multiply,
			3: save,
			4: printFunc,
			5: jumpIfTrue,
			6: jumpIfFalse,
			7: lessThan,
			8: equals,
		},
	}
	//computer.Run([]int{3, 0, 4, 0, 99})

	programString := adventutil.Parse(5)[0]
	partsStrings := strings.Split(programString, ",")
	var program []int
	for _, partString := range partsStrings {
		asInt, _ := strconv.Atoi(partString)
		program = append(program, asInt)
	}

	computer.Run(program)
}
