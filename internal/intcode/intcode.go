package intcode

import (
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"strconv"
	"strings"
)

type IPType int64
type memory map[IPType]IPType
type Instruction func(Computer, memory, IPType) IPType

func ParseProgram(programString string) []IPType {
	partsStrings := strings.Split(programString, ",")
	var program []IPType
	for _, partString := range partsStrings {
		asInt, _ := strconv.ParseInt(partString, 10, 64)
		program = append(program, IPType(asInt))
	}
	return program
}

type Computer struct {
	name    string
	ops     map[int]Instruction
	Input   chan IPType
	Output  chan IPType
	Stopped chan struct{}
}

func CreateComputer(name string, ops map[int]Instruction) Computer {
	return Computer{name, ops, make(chan IPType), make(chan IPType), make(chan struct{})}
}

func binaryFunc(program memory, ip IPType, operand func(IPType, IPType) IPType) IPType {
	arg1 := getValue(program, ip, 1)
	arg2 := getValue(program, ip, 2)
	dest := program[ip+3]
	program[dest] = operand(arg1, arg2)
	return ip + 4
}

func Add(c Computer, program memory, ip IPType) IPType {
	return binaryFunc(program, ip, func(i IPType, i2 IPType) IPType {
		return i + i2
	})
}

func Multiply(c Computer, program memory, ip IPType) IPType {
	return binaryFunc(program, ip, func(i IPType, i2 IPType) IPType {
		return i * i2
	})
}

func Save(c Computer, program memory, ip IPType) IPType {
	//fmt.Println(program[ip:ip+2])
	dest := program[ip+1]
	input := <-c.Input
	//fmt.Printf("%s received input %d\n", c.name, input)
	program[dest] = input
	return ip + 2
}

func PrintFunc(c Computer, program memory, ip IPType) IPType {
	value := getValue(program, ip, 1)
	//fmt.Printf("%s output %d\n", c.name, value)
	c.Output <- value
	return ip + 2
}

func jump(program memory, ip IPType, shouldJump func(IPType) bool) IPType {
	arg1 := getValue(program, ip, 1)
	if shouldJump(arg1) {
		return getValue(program, ip, 2)
	}
	return ip + 3
}

func JumpIfTrue(c Computer, program memory, ip IPType) IPType {
	return jump(program, ip, func(i IPType) bool {
		return i != 0
	})
}

func JumpIfFalse(c Computer, program memory, ip IPType) IPType {
	return jump(program, ip, func(i IPType) bool {
		return i == 0
	})
}

func LessThan(c Computer, program memory, ip IPType) IPType {
	return binaryFunc(program, ip, func(i IPType, i2 IPType) IPType {
		if i < i2 {
			return 1
		}
		return 0
	})
}

func Equals(c Computer, program memory, ip IPType) IPType {
	return binaryFunc(program, ip, func(i IPType, i2 IPType) IPType {
		if i == i2 {
			return 1
		}
		return 0
	})
}

func getValue(program memory, ip IPType, argNum int) IPType {
	opcode := program[ip]
	argValue := program[ip+IPType(argNum)]
	if isImmediateMode(opcode, argNum) {
		return argValue
	}
	return program[argValue]
}

func isImmediateMode(opcode IPType, argNum int) bool {
	return (opcode/(IPType(adventutil.Pow10(argNum+1))))%10 == 1
}

func (c Computer) Run(program []IPType) {
	memory := make(memory)
	for i, val := range program {
		memory[IPType(i)] = IPType(val)
	}

	ip := IPType(0)
	for {
		opcode := int(memory[ip] % 100)
		//fmt.Printf("%s IP %d op %d %v\n", c.name, ip, executable[ip], executable)
		if opcode == 99 {
			//fmt.Println(c.name + " ending")
			close(c.Output)
			close(c.Stopped)
			break
		}
		ip = c.ops[opcode](c, memory, ip)
	}
}
