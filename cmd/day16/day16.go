package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"strconv"
)

func ApplyPhase(input []int, offset int) []int {
	//fmt.Println("Applying phase")
	totalLen := len(input)
	var output []int
	for x :=0; x < offset; x++ {
		output = append(output, 0)
	}
	for outputDigit := offset; outputDigit < totalLen; outputDigit++ {
		//fmt.Printf("Digit %d = \n", outputDigit)
		outputValue := 0
		for i := outputDigit; i <= (outputDigit*2) && i < totalLen; i++ {
			for j := i; j < totalLen; j += 4 * (outputDigit + 1) {
				//fmt.Printf("+%d (input[%d])\n", input[j], j)
				outputValue += input[j]
			}
		}
		subtractStart := 2 + 3*outputDigit
		for i := subtractStart; i <= subtractStart+outputDigit && i < totalLen; i++ {
			for j := i; j < totalLen; j += 4 * (outputDigit + 1) {
				//fmt.Printf("-%d (input[%d])\n", input[j], j)
				outputValue -= input[j]
			}
		}
		//fmt.Printf("total = %d\n", outputValue)
		outputValue = adventutil.Abs(outputValue) % 10
		//fmt.Printf(" = %d\n", outputValue)
		output = append(output, outputValue)
	}
	return output
}

func ApplyPhase2(input []int, offset int) []int {
	if offset < len(input) / 2 {
		panic("Assumption violated")
	}

	result := make([]int, len(input))

	sum := 0
	for outputDigit := offset; outputDigit < len(input); outputDigit++ {
		sum += input[outputDigit]
	}
	for outputDigit := offset; outputDigit < len(input); outputDigit++ {
		result[outputDigit] = adventutil.Abs(sum) % 10
		sum -= input[outputDigit]
	}
	return result
}

func Apply100(val []int, offset int, phaseFunc func([]int, int) []int) []int {
	for i := 0; i < 100; i++ {
		val = phaseFunc(val, offset)
	}
	return val
}

func parseSignal(input string) []int {
	var output []int
	for _, digitChar := range input {
		asInt, _ := strconv.Atoi(string(digitChar))
		output = append(output, asInt)
	}
	return output
}

func FullFFT(input string) string {
	baseInputArray := parseSignal(input)
	var inputArray []int
	for i := 0; i < 10000; i++ {
		inputArray = append(inputArray, baseInputArray...)
	}
	offset, _ := strconv.Atoi(input[:7])
	result := Apply100(inputArray, offset, ApplyPhase2)
	return DigitsAtOffset(result, offset)
}

func DigitsAtOffset(values []int, offset int) string {
	var result string
	for i := 0; i < 8; i++ {
		result += fmt.Sprintf("%d", values[i+offset])
	}
	return result
}

func main() {
	inputStr := adventutil.Parse(16)[0]
	val := parseSignal(inputStr)
	result := Apply100(val, 0, ApplyPhase)
	fmt.Printf("Part 1: %s\n", DigitsAtOffset(result, 0))

	part2 := FullFFT(inputStr)
	fmt.Printf("Part 2: %s\n", part2)
}
