package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	. "github.com/enjean/advent-of-code-2019/internal/adventutil/coordinate"
	. "github.com/enjean/advent-of-code-2019/internal/intcode"
)

func BuildMap(program Program) map[Coordinate]int {
	computer := CreateCompleteComputer("ASCII")
	go func() { computer.Run(program) }()

	shipMap := make(map[Coordinate]int)

	currentPos := Coordinate{}

	for output := range computer.Output {
		//fmt.Printf("%s = %d\n", currentPos, output)
		if output == 10 {
			currentPos = Coordinate{Y: currentPos.Y + 1}
			continue
		}
		shipMap[currentPos] = int(output)
		currentPos = Coordinate{X: currentPos.X + 1, Y: currentPos.Y}
	}

	return shipMap
}

func FindIntersections(shipMap map[Coordinate]int) []Coordinate {
	var intersections []Coordinate
	for coord, val := range shipMap {
		if val != 35 {
			continue
		}
		if isIntersection(shipMap, coord) {
			intersections = append(intersections, coord)
		}
	}
	return intersections
}

func isIntersection(shipMap map[Coordinate]int, coordinate Coordinate) bool {
	for _, neighbor := range coordinate.Neighbors() {
		if shipMap[neighbor] != 35 {
			return false
		}
	}
	return true
}

func main() {
	program := ParseProgram(adventutil.Parse(17)[0])
	shipMap := BuildMap(program)
	PrintIntCoordinateMap(shipMap, func(i int) string {
		return string(i)
	})

	intersections := FindIntersections(shipMap)
	alignmentSum := 0
	for _, intersection := range intersections {
		alignmentSum += intersection.X * intersection.Y
	}
	fmt.Printf("Part 1: %d\n", alignmentSum)
}
