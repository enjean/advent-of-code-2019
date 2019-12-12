package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	. "github.com/enjean/advent-of-code-2019/internal/adventutil/coordinate"
	. "github.com/enjean/advent-of-code-2019/internal/intcode"
)

type Direction int

const (
	U Direction = iota + 1
	D
	L
	R
)

func PaintPanels(program []IPType, startingColor int) (map[Coordinate]bool, map[Coordinate]int) {
	colors := make(map[Coordinate]int)
	painted := make(map[Coordinate]bool)

	computer := CreateComputer("Painter", map[int]Instruction{
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

	currentCoord := Coordinate{}
	colors[currentCoord] = startingColor

	currentDir := U
	for {
		paintType, done := getOutput(computer, IPType(colors[currentCoord]))
		if done {
			break
		}
		colors[currentCoord] = int(paintType)
		painted[currentCoord] = true

		turnInstruction, done := getOutput(computer, IPType(colors[currentCoord]))
		if done {
			break
		}
		currentDir = turn(currentDir, int(turnInstruction))
		currentCoord = move(currentCoord, currentDir)
	}
	return painted, colors
}

func getOutput(computer *Computer, input IPType) (IPType, bool) {
	for {
		select {
		case <-computer.Stopped:
			return -1, true
		case computer.Input <- input:
		case output := <-computer.Output:
			return output, false
		}
	}
}

func turn(currentDirection Direction, turnInstruction int) Direction {
	switch currentDirection {
	case U:
		if turnInstruction == 0 {
			return L
		}
		return R
	case R:
		if turnInstruction == 0 {
			return U
		}
		return D
	case D:
		if turnInstruction == 0 {
			return R
		}
		return L
	case L:
		if turnInstruction == 0 {
			return D
		}
		return U
	}
	panic("Invalid direction")
}

func move(position Coordinate, direction Direction) Coordinate {
	switch direction {
	case U:
		return Coordinate{X: position.X, Y: position.Y - 1}
	case R:
		return Coordinate{X: position.X + 1, Y: position.Y}
	case D:
		return Coordinate{X: position.X, Y: position.Y + 1}
	case L:
		return Coordinate{X: position.X - 1, Y: position.Y}
	}
	panic("Invalid direction")
}

func main() {
	program := ParseProgram(adventutil.Parse(11)[0])
	part1Painted, _ := PaintPanels(program, 0)
	fmt.Printf("Part 1: %d\n", len(part1Painted))

	_, part2Colors := PaintPanels(program, 1)

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	for c := range part2Colors {
		if c.X < minX {
			minX = c.X
		}
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}
	for y := minX; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			var character string
			coord := Coordinate{X: x, Y: y}
			if part2Colors[coord] == 1 {
				character = "*"
			} else {
				character = " "
			}
			fmt.Print(character)
		}
		fmt.Println()
	}

}
