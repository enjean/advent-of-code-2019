package intcode

import (
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"strconv"
	"strings"
)

func ParseProgram(programString string) []int {
	partsStrings := strings.Split(programString, ",")
	var program []int
	for _, partString := range partsStrings {
		asInt, _ := strconv.Atoi(partString)
		program = append(program, asInt)
	}
	return program
}

type Computer struct {
	name    string
	ops     map[int]func(Computer, []int, int) int
	Input   chan int
	Output  chan int
	Stopped chan struct{}
}

func CreateComputer(name string, ops map[int]func(Computer, []int, int) int) Computer {
	return Computer{name, ops, make(chan int), make(chan int), make(chan struct{})}
}

func binaryFunc(program []int, ip int, operand func(int, int) int) int {
	arg1 := getValue(program, ip, 1)
	arg2 := getValue(program, ip, 2)
	dest := program[ip+3]
	program[dest] = operand(arg1, arg2)
	return ip + 4
}

func Add(c Computer, program []int, ip int) int {
	return binaryFunc(program, ip, func(i int, i2 int) int {
		return i + i2
	})
}

func Multiply(c Computer, program []int, ip int) int {
	return binaryFunc(program, ip, func(i int, i2 int) int {
		return i * i2
	})
}

func Save(c Computer, program []int, ip int) int {
	//fmt.Println(program[ip:ip+2])
	dest := program[ip+1]
	input := <-c.Input
	//fmt.Printf("%s received input %d\n", c.name, input)
	program[dest] = input
	return ip + 2
}

func PrintFunc(c Computer, program []int, ip int) int {
	value := getValue(program, ip, 1)
	//fmt.Printf("%s output %d\n", c.name, value)
	c.Output <- value
	return ip + 2
}

func jump(program []int, ip int, shouldJump func(int) bool) int {
	arg1 := getValue(program, ip, 1)
	if shouldJump(arg1) {
		return getValue(program, ip, 2)
	}
	return ip + 3
}

func JumpIfTrue(c Computer, program []int, ip int) int {
	return jump(program, ip, func(i int) bool {
		return i != 0
	})
}

func JumpIfFalse(c Computer, program []int, ip int) int {
	return jump(program, ip, func(i int) bool {
		return i == 0
	})
}

func LessThan(c Computer, program []int, ip int) int {
	return binaryFunc(program, ip, func(i int, i2 int) int {
		if i < i2 {
			return 1
		}
		return 0
	})
}

func Equals(c Computer, program []int, ip int) int {
	return binaryFunc(program, ip, func(i int, i2 int) int {
		if i == i2 {
			return 1
		}
		return 0
	})
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
	executable := make([]int, len(program))
	copy(executable, program)

	ip := 0
	for {
		opcode := executable[ip] % 100
		//fmt.Printf("%s IP %d op %d %v\n", c.name, ip, executable[ip], executable)
		if opcode == 99 {
			//fmt.Println(c.name + " ending")
			close(c.Output)
			close(c.Stopped)
			break
		}
		ip = c.ops[opcode](c, executable, ip)
	}
}
