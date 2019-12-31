package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	. "github.com/enjean/advent-of-code-2019/internal/adventutil/coordinate"
	. "github.com/enjean/advent-of-code-2019/internal/intcode"
	"strconv"
)

type Direction int

const (
	N Direction = 1
	S Direction = 2
	W Direction = 3
	E Direction = 4
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

func findPath(shipMap map[Coordinate]int) []string {
	var currentLocation Coordinate
	var currentDirection Direction
	for coord, val := range shipMap {
		switch val {
		case '^':
			currentLocation = coord
			currentDirection = N
		case '>':
			currentLocation = coord
			currentDirection = E
		case 'v':
			currentLocation = coord
			currentDirection = S
		case '<':
			currentLocation = coord
			currentDirection = W
		}
	}

	visited := make(map[Coordinate]int)

	var path []string
	stepCounter := 0
	for {
		visited[currentLocation] = 1
		straightNeighbor := neighborInDirection(currentLocation, currentDirection)
		if shipMap[straightNeighbor] == 35 {
			currentLocation = straightNeighbor
			stepCounter++
			continue
		}

		leftTurnDir := leftTurnDir(currentDirection)
		leftNeighbor := neighborInDirection(currentLocation, leftTurnDir)
		if shipMap[leftNeighbor] == 35 {
			if stepCounter > 0 {
				path = append(path, strconv.Itoa(stepCounter))
				stepCounter = 0
			}
			path = append(path, "L")
			currentDirection = leftTurnDir
			continue
		}
		rightTurnDir := rightTurnDir(currentDirection)
		rightNeighbor := neighborInDirection(currentLocation, rightTurnDir)
		if shipMap[rightNeighbor] == 35 {
			if stepCounter > 0 {
				path = append(path, strconv.Itoa(stepCounter))
				stepCounter = 0
			}
			path = append(path, "R")
			currentDirection = rightTurnDir
			continue
		}

		// reached end of path
		if stepCounter > 0 {
			path = append(path, strconv.Itoa(stepCounter))
			stepCounter = 0
		}
		break
	}

	PrintIntCoordinateMap(visited, func(i int) string {
		if i == 1 {
			return "X"
		}
		return "."
	})

	return path
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

func leftTurnDir(direction Direction) Direction {
	switch direction {
	case N:
		return W
	case E:
		return N
	case S:
		return E
	case W:
		return S
	}
	panic("Bad direction")
}

func rightTurnDir(direction Direction) Direction {
	switch direction {
	case N:
		return E
	case E:
		return S
	case S:
		return W
	case W:
		return N
	}
	panic("Bad direction")
}

func runVacuum(mainRoutine, movementStringA, movementStringB, movementStringC string, program Program) {
	program[0] = 2
	computer := CreateCompleteComputer("ASCII")
	go func() { computer.Run(program) }()

	//var lastVal IPType
	for {
		line, _ := readLine(computer)
		fmt.Println(line)
		if line == "Main:" {
			inputLine(mainRoutine, computer)
		}
		if line == "Function A:" {
			inputLine(movementStringA, computer)
		}
		if line == "Function B:" {
			inputLine(movementStringB, computer)
		}
		if line == "Function C:" {
			inputLine(movementStringC, computer)
		}
		if line == "Continuous video feed?" {
			inputLine("n", computer)
			break
		}
		//if val == 10 && lastVal == 10 {
		//	break
		//}
		//lastVal = val
	}

	//go func() {
	//	for output := range computer.Output {
	//		fmt.Printf("Output %d = %s\n", output, string(output))
	//	}
	//}()
	for {
		line, done := readLine(computer)
		fmt.Println(line)
		if done {
			break
		}
	}

}

func readLine(computer *Computer) (string, bool) {
	var line string
	for {
		select {
		case val := <-computer.Output:
			if val == 10 {
				return line, false
			}
			if val > 255 {
				line += fmt.Sprintf("%d", val)
				continue
			}
			line += string(val)
			case <- computer.Stopped:
				return line, true
		}
	}
}

func inputLine(line string, computer *Computer) {
	for _, ascii := range line {
		computer.Input <- IPType(ascii)
	}
	computer.Input <- IPType(10)
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

	path := findPath(shipMap)
	fmt.Printf("Path = %v\n", path)

	//L 12 L 12 R 12
	//L 12 L 12 R 12
	//L 8 L 8 R 12 L 8 L 8
	//L 10 R 8 R 12
	//L 10 R 8 R 12
	//L 12 L 12 R 12
	//L 8 L 8 R 12 L 8 L 8
	//L 10 R 8 R 12
	//L 12 L 12 R 12
	//L 8 L 8 R 12 L 8 L 8
	movementStringA := "L,12,L,12,R,12"
	movementStringB := "L,8,L,8,R,12,L,8,L,8"
	movementStringC := "L,10,R,8,R,12"
	mainRoutine := "A,A,B,C,C,A,B,C,A,B"

	runVacuum(mainRoutine, movementStringA, movementStringB, movementStringC, program)
}
