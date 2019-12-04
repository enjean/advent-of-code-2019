package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/adventutil"
	"strconv"
	"strings"
)

type Direction int

const (
	U Direction = iota + 1
	D
	L
	R
)

type PathComponent struct {
	dir      Direction
	distance int
}

type Wire []PathComponent

type point struct {
	x, y int
}

func (p point) distance(o point) int {
	return adventutil.Abs(p.x-o.x) + adventutil.Abs(p.y-o.y)
}

func FindNearestIntersection(wires []Wire) int {
	wiresAtPoints := make(map[point]int)

	for _, wire := range wires {
		pointsInWire := make(map[point]bool)
		currentPosition := point{0, 0}
		for _, pathComponent := range wire {
			movementFunc := movementFunction(pathComponent.dir)
			for i := 0; i < pathComponent.distance; i++ {
				currentPosition = movementFunc(currentPosition)
				pointsInWire[currentPosition] = true
			}
		}
		for p := range pointsInWire {
			wiresAtPoints[p]++
		}
	}

	minDistance := adventutil.MaxInt
	for p, numWires := range wiresAtPoints {
		if numWires == 1 {
			continue
		}
		distanceToIntersection := point{0, 0}.distance(p)
		if distanceToIntersection < minDistance {
			minDistance = distanceToIntersection
		}
	}

	return minDistance
}

func FindFirstIntersection(wire1, wire2 Wire) int {
	wire1StepsToPoint := buildStepsToPoint(wire1)
	wire2StepsToPoint := buildStepsToPoint(wire2)

	minSteps := adventutil.MaxInt
	for p, count := range wire1StepsToPoint {
		w2Count, ok := wire2StepsToPoint[p]
		if ok {
			if count+w2Count < minSteps {
				minSteps = count + w2Count
			}
		}
	}

	return minSteps
}

func buildStepsToPoint(wire Wire) map[point]int {
	wirePoints := make(map[point]int)
	step := 0
	point := point{0, 0}
	for _, pc := range wire {
		for i := 0; i < pc.distance; i++ {
			step++
			point = movementFunction(pc.dir)(point)
			if wirePoints[point] == 0 {
				wirePoints[point] = step
			}
		}
	}
	return wirePoints
}

func movementFunction(direction Direction) func(point) point {
	switch direction {
	case U:
		return func(p point) point {
			return point{p.x, p.y + 1}
		}
	case D:
		return func(p point) point {
			return point{p.x, p.y - 1}
		}
	case L:
		return func(p point) point {
			return point{p.x - 1, p.y}
		}
	case R:
		return func(p point) point {
			return point{p.x + 1, p.y}
		}
	}
	panic("Invalid dir")
}

func main() {
	lines := adventutil.Parse(3)

	var wires []Wire
	for _, line := range lines {
		lineParts := strings.Split(line, ",")
		var wire Wire
		for _, pathComponentString := range lineParts {
			dirString := string(pathComponentString[0])
			var direction Direction
			switch dirString {
			case "U":
				direction = U
			case "D":
				direction = D
			case "L":
				direction = L
			case "R":
				direction = R
			}
			distanceString := pathComponentString[1:]
			distance, _ := strconv.Atoi(distanceString)
			wire = append(wire, PathComponent{direction, distance})
		}
		wires = append(wires, wire)
	}

	distanceToClosestIntersection := FindNearestIntersection(wires)
	fmt.Printf("Part 1: %d\n", distanceToClosestIntersection)

	part2 := FindFirstIntersection(wires[0], wires[1])
	fmt.Printf("Part 2: %d\n", part2)
}
