package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"strconv"
)

var pattern = [4]int{0, 1, 0, -1}

func ApplyPhase(input []int) []int {
	var output []int
	for outputDigit := 0; outputDigit < len(input); outputDigit++ {
		outputValue := 0
		for i, val := range input {
			patternIndex := ((i + 1) / (outputDigit + 1)) % 4
			//fmt.Printf("%d*%d + ", val, pattern[patternIndex])
			outputValue += val * pattern[patternIndex]
		}
		outputValue = adventutil.Abs(outputValue) % 10
		//fmt.Printf(" = %d\n", outputValue)
		output = append(output, outputValue)
	}
	return output
}

func parseSignal(input string) []int {
	var output []int
	for _, digitChar := range input {
		asInt, _ := strconv.Atoi(string(digitChar))
		output = append(output, asInt)
	}
	return output
}

func main() {
	val := parseSignal(adventutil.Parse(16)[0])
	for i := 0; i < 100; i++ {
		val = ApplyPhase(val)
	}

	fmt.Printf("Part 1: ")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d", val[i])
	}
	fmt.Println()
}

