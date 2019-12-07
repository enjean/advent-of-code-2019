package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"strings"
)

func CountOrbits(orbitStrings []string) int {
	orbitsMap := toOrbitsMap(orbitStrings)
	orbits := 0
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

func OrbitalTransfersNeeded(orbitStrings []string, from, to string) int {
	orbitsMap := toOrbitsMap(orbitStrings)
	lca := lowestCommonAncestor(orbitsMap, from, to)
	fmt.Println("LCA is " + lca)

	distance1 := distance(orbitsMap, orbitsMap[from], lca)
	fmt.Printf("Orbital transfers from %s to %s = %d\n", from, lca, distance1)
	distance2 := distance(orbitsMap, orbitsMap[to], lca)
	fmt.Printf("Orbital transfers from %s to %s = %d\n", to, lca, distance2)

	return distance1 + distance2
}

func toOrbitsMap(orbitStrings []string) map[string]string {
	orbitsMap := make(map[string]string)
	for _, orbitString := range orbitStrings {
		parts := strings.Split(orbitString, ")")
		target := parts[0]
		orbitedBy := parts[1]
		orbitsMap[orbitedBy] = target
	}
	return orbitsMap
}

func lowestCommonAncestor(orbitsMap map[string]string, from, to string) string {
	pathToRootForFrom := make(map[string]bool)
	objectName := from
	for {
		pathToRootForFrom[objectName] = true
		if objectName == "COM" {
			break
		}
		objectName = orbitsMap[objectName]
	}

	objectName = to
	for {
		if pathToRootForFrom[objectName] {
			break
		}
		objectName = orbitsMap[objectName]
	}

	return objectName
}

func distance(orbitsMap map[string]string, from, to string) int {
	result := 0

	node := from
	for {
		if node == to {
			break
		}
		node = orbitsMap[node]
		result++
	}

	return result
}

func main() {
	orbitStrings := adventutil.Parse(6)

	numOrbits := CountOrbits(orbitStrings)
	fmt.Printf("Part 1: %d\n", numOrbits)

	orbitalTransfers := OrbitalTransfersNeeded(orbitStrings, "YOU", "SAN")
	fmt.Printf("Part 12 %d\n", orbitalTransfers)
}
