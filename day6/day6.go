package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/adventutil"
	"strings"
)

//type spaceObject struct {
//	label string
//	orbitedBy *spaceObject
//}

// B -> COM
// C -> B
func CountOrbits(orbitStrings []string) int {
	orbits := 0
	orbitsMap := make(map[string]string)
	for _, orbitString := range orbitStrings {
		parts := strings.Split(orbitString,")")
		target := parts[0]
		orbitedBy := parts[1]
		orbitsMap[orbitedBy] = target
	}
	for _, target := range orbitsMap {
		for {
			orbits++
			if target == "COM" {
				break
			}
			target = orbitsMap[target]
		}
	}
	return orbits
}

func main() {
	orbitStrings := adventutil.Parse(6)

	numOrbits := CountOrbits(orbitStrings)
	fmt.Printf("Part 1: %d\n", numOrbits)
}
