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
	for i := 0; i < len(pieces)-1; i++ {
		if pieces[i] == pieces[i+1] {
			hasDouble = true
		}
		if pieces[i] > pieces[i+1] {
			return false
		}
	}
	return hasDouble
}

func NumValid(start, end int, validateFunc func(string) bool) int {
	numValid := 0
	for i := start; i <= end; i++ {
		if validateFunc(strconv.Itoa(i)) {
			numValid++
		}
	}
	return numValid
}

func Validate2(password string) bool {
	if !Validate(password) {
		return false
	}
	pieces := []rune(password)
	groupSizes := make(map[rune]int)
	for _, piece := range pieces {
		groupSizes[piece]++
	}
	for _, count := range groupSizes {
		if count == 2 {
			return true
		}
	}
	return false
}

func main() {
	start := 356261
	end := 846303

	part1 := NumValid(start, end, Validate)
	fmt.Printf("Part 1: %d \n", part1)

	part2 := NumValid(start, end, Validate2)
	fmt.Printf("Part 2: %d \n", part2)
}
