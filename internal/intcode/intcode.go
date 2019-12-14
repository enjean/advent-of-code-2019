package intcode

import (
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"strconv"
	"strings"
)

type IPType int64
type memory map[IPType]IPType
type Instruction func(*Computer, memory, IPType) IPType

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
	name         string
	ops          map[int]Instruction
	Input        chan IPType
	Output       chan IPType
	Stopped      chan struct{}
	relativeBase IPType
}

func CreateComputer(name string, ops map[int]Instruction) *Computer {
	return &Computer{name, ops, make(chan IPType), make(chan IPType), make(chan struct{}), 0}
}

func binaryFunc(c *Computer, program memory, ip IPType, operand func(IPType, IPType) IPType) IPType {
	//fmt.Printf("%d %d %d %d\n", program[ip], program[ip+1],program[ip+2],program[ip+3])
	arg1 := getValue(c, program, ip, 1)
	arg2 := getValue(c, program, ip, 2)
	dest := getWriteAddress(c, program, ip, 3)
	//fmt.Printf("program[%d] = %d op %d = %d\n", dest, arg1, arg2,  operand(arg1, arg2))
	program[dest] = operand(arg1, arg2)
	return ip + 4
}

func Add(c *Computer, program memory, ip IPType) IPType {
	return binaryFunc(c, program, ip, func(i IPType, i2 IPType) IPType {
		return i + i2
	})
}

func Multiply(c *Computer, program memory, ip IPType) IPType {
	return binaryFunc(c, program, ip, func(i IPType, i2 IPType) IPType {
		return i * i2
	})
}

func Save(c *Computer, program memory, ip IPType) IPType {
	//fmt.Printf("save %d %d\n", program[ip], program[ip+1])
	input := <-c.Input
	dest := getWriteAddress(c, program, ip, 1)
	//fmt.Printf("%s received input %d\n", c.name, input)
	//fmt.Printf("program[%d]=%d\n", dest, input)
	program[dest] = input
	return ip + 2
}

func PrintFunc(c *Computer, program memory, ip IPType) IPType {
	value := getValue(c, program, ip, 1)
	//fmt.Printf("%s output %d\n", c.name, value)
	c.Output <- value
	return ip + 2
}

func jump(c *Computer, program memory, ip IPType, shouldJump func(IPType) bool) IPType {
	//fmt.Printf("%d %d %d \n", program[ip], program[ip+1],program[ip+2])
	arg1 := getValue(c, program, ip, 1)
	if shouldJump(arg1) {
		return getValue(c, program, ip, 2)
	}
	return ip + 3
}

func JumpIfTrue(c *Computer, program memory, ip IPType) IPType {
	return jump(c, program, ip, func(i IPType) bool {
		return i != 0
	})
}

func JumpIfFalse(c *Computer, program memory, ip IPType) IPType {
	return jump(c, program, ip, func(i IPType) bool {
		return i == 0
	})
}

func LessThan(c *Computer, program memory, ip IPType) IPType {
	return binaryFunc(c, program, ip, func(i IPType, i2 IPType) IPType {
		if i < i2 {
			return 1
		}
		return 0
	})
}

func Equals(c *Computer, program memory, ip IPType) IPType {
	return binaryFunc(c, program, ip, func(i IPType, i2 IPType) IPType {
		if i == i2 {
			return 1
		}
		return 0
	})
}

func AdjustRelativeBase(c *Computer, program memory, ip IPType) IPType {
	c.relativeBase += getValue(c, program, ip, 1)
	return ip + 2
}

func getWriteAddress(c *Computer, program memory, ip IPType, argNum int) IPType {
	opcode := program[ip]
	argValue := program[ip+IPType(argNum)]
	mode := parameterMode(opcode, argNum)
	if mode == 2 {
		argValue += c.relativeBase
	}
	return argValue
}

func getValue(c *Computer, program memory, ip IPType, argNum int) IPType {
	opcode := program[ip]
	argValue := program[ip+IPType(argNum)]
	mode := parameterMode(opcode, argNum)
	if mode == 0 {
		if argValue < 0 {
			panic("Invalid index")
		}
		//fmt.Printf("Returning program[%d]=%d\n", argValue, program[argValue])
		return program[argValue]
	}
	if mode == 1 {
		return argValue
	}
	if mode == 2 {
		index := argValue + c.relativeBase
		if index < 0 {
			panic("Invalid index")
		}
		return program[index]
	}
	panic("Unknown mode")
}

func parameterMode(opcode IPType, argNum int) int {
	return int((int64(opcode) / (adventutil.Pow10(argNum + 1))) % 10)
}

func (c *Computer) Run(program []IPType) {
	memory := make(memory)
	for i, val := range program {
		memory[IPType(i)] = val
	}

	ip := IPType(0)
	for {
		opcode := int(memory[ip] % 100)
//		fmt.Printf("%s IP %d op %d  rb %d %v\n", c.name, ip, memory[ip], c.relativeBase, memory)
		if opcode == 99 {
			//fmt.Println(c.name + " ending")
			close(c.Output)
			close(c.Stopped)
			break
		}
		ip = c.ops[opcode](c, memory, ip)
	}
}
