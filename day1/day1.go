package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/adventutil"
	"strconv"
)

func FuelNeeded(mass int64) int64 {
	return mass / 3 - 2
}

func main() {
	lines := adventutil.Parse(1)
	var sum int64
	for _, line := range lines {
		mass, _ := strconv.ParseInt(line, 10, 64)
		sum += FuelNeeded(int64(mass))
	}
	fmt.Printf("Part 1: Sum = %d", sum)
}
