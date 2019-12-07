package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"strconv"
)

func FuelNeeded(mass int64) int64 {
	return mass / 3 - 2
}

func CumulativeFuelNeeded(mass int64) int64 {
	var total int64
	massToCalc := mass
	for true {
		fuelNeeded := FuelNeeded(massToCalc)
		if fuelNeeded <= 0 {
			break
		}
		total += fuelNeeded
		massToCalc = fuelNeeded
	}
	return total
}

func main() {
	lines := adventutil.Parse(1)
	var sum int64
	var cumulativeFuelNeeded int64
	for _, line := range lines {
		mass, _ := strconv.ParseInt(line, 10, 64)
		sum += FuelNeeded(mass)
		cumulativeFuelNeeded += CumulativeFuelNeeded(mass)
	}
	fmt.Printf("Part 1: Sum = %d\n", sum)
	fmt.Printf("Part 2: Cumulative Fuel Needed = %d\n", cumulativeFuelNeeded)
}
