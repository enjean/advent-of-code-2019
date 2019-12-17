package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	. "github.com/enjean/advent-of-code-2019/internal/adventutil/coordinate"
	. "github.com/enjean/advent-of-code-2019/internal/intcode"
)

type Direction int

const (
	N Direction = 1
	S Direction = 2
	W Direction = 3
	E Direction = 4
)

func directions() []Direction {
	return []Direction{N,S,W,E}
}

func (d Direction) reverse() Direction {
	switch d {
	case N: return S
	case S: return N
	case W: return E
	case E: return W
	}
	panic("Bad direction")
}

func neighborInDirection(coordinate Coordinate, direction Direction) Coordinate {
	switch direction {
	case N:
		return Coordinate{X: coordinate.X, Y: coordinate.Y - 1}
	case S:
		return Coordinate{X: coordinate.X, Y: coordinate.Y + 1}
	case W:
		return Coordinate{X: coordinate.X - 1, Y: coordinate.Y}
	case E:
		return Coordinate{X: coordinate.X + 1, Y: coordinate.Y}
	}
	panic("Bad direction")
}

type directionCoord struct {
		direction Direction
		from      Coordinate
}

func ShortestPathToOxygenSystem(program []IPType) int {
	computer := CreateComputer("Game", map[int]Instruction{
		1: Add,
		2: Multiply,
		3: Save,
		4: PrintFunc,
		5: JumpIfTrue,
		6: JumpIfFalse,
		7: LessThan,
		8: Equals,
		9: AdjustRelativeBase,
	})
	go func() { computer.Run(program) }()

	visited := make(map[Coordinate]bool)
	stepsTo := make(map[Coordinate]int)
	dirToOrigin := make(map[Coordinate]Direction)
	var coordsToTry []directionCoord
	origin := Coordinate{0,0}
	currentLocation := Coordinate{}
	stepsTo[currentLocation] = 0
	for _, direction := range directions() {
		coordsToTry = append(coordsToTry, directionCoord{direction:direction, from:currentLocation})
	}
	for {
		trying := coordsToTry[0]
		fmt.Printf("Trying %v from %v \n", trying.direction, trying.from)

		coordsToTry = coordsToTry[1:]

		sourceCoord := trying.from
		// travel there
		if currentLocation != sourceCoord {
			fmt.Printf("Need to get to %v from %v\n", sourceCoord, currentLocation)
			currentLocToOrigin := []Coordinate{currentLocation}
			inPath := map[Coordinate]bool{currentLocation:true}
			loc := currentLocation
			for loc != origin {
				loc = neighborInDirection(loc, dirToOrigin[loc])
				currentLocToOrigin = append(currentLocToOrigin, loc)
				inPath[loc] = true
			}

			var commonAncestor Coordinate
			var directionsLCAToSource []Direction
			loc = sourceCoord
			for {
				if inPath[loc] {
					commonAncestor = loc
					break
				}
				direction := dirToOrigin[loc]
				directionsLCAToSource = append([]Direction{direction.reverse()}, directionsLCAToSource...)
				loc = neighborInDirection(loc, dirToOrigin[loc])
			}

			for currentLocation != commonAncestor {
				computer.Input <- IPType(dirToOrigin[currentLocation])
				if output := <-computer.Output; output != 1 {
					panic("Unexpected")
				}
				currentLocation = neighborInDirection(currentLocation, dirToOrigin[currentLocation])
			}
			for _, dir := range directionsLCAToSource {
				computer.Input <- IPType(dir)
				if output := <-computer.Output; output != 1 {
					panic("Unexpected")
				}
				currentLocation = neighborInDirection(currentLocation, dir)
			}
			if currentLocation != sourceCoord {
				panic("Not at sourceCoord")
			}
		}

		destCoord := neighborInDirection(sourceCoord, trying.direction)
		visited[destCoord] = true
		computer.Input <- IPType(trying.direction)

		result := <-computer.Output
		if result == 2 {
			return stepsTo[sourceCoord] + 1
		}
		if result == 0 {
			fmt.Printf("%v = wall\n", destCoord)
			// hit a wall
			continue
		}
		if result == 1 {
			currentLocation = destCoord
			stepsTo[currentLocation] = stepsTo[sourceCoord] + 1
			dirToOrigin[currentLocation] = trying.direction.reverse()
			for _, direction := range directions() {
				neighbor := neighborInDirection(destCoord, direction)
				if !visited[neighbor] {
					coordsToTry = append(coordsToTry, directionCoord{direction:direction, from:currentLocation})
				}
			}
			// moved there
		}
	}
}

func main() {
	part1 := ShortestPathToOxygenSystem(ParseProgram(adventutil.Parse(15)[0]))
	fmt.Printf("Part 1: %d\n", part1)
}
