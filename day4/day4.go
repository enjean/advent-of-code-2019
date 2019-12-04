package main

import (
	"fmt"
	"strconv"
)

//It is a six-digit number.
//The value is within the range given in your puzzle input.
//Two adjacent digits are the same (like 22 in 122345).
//Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).

func Validate(password string) bool {
	pieces := []rune(password)
	hasDouble := false
	for i := 0; i < len(pieces) - 1; i++ {
		if pieces[i] == pieces[i + 1] {
			hasDouble = true
		}
		if pieces[i] > pieces[i+1] {
			return false
		}
	}
	return hasDouble
}

func NumValid(start, end int) int {
	numValid := 0
	for i := start; i <= end; i++ {
		if Validate(strconv.Itoa(i)) {
			numValid++
		}
	}
	return numValid
}

func main() {
	start := 356261
	end := 846303

	part1 := NumValid(start, end)
	fmt.Printf("Part 1: %d \n", part1)
}
