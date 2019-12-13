package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	. "github.com/enjean/advent-of-code-2019/internal/adventutil/coordinate"
	. "github.com/enjean/advent-of-code-2019/internal/intcode"
)

func buildGrid(program []IPType) map[Coordinate]int {
	grid := make(map[Coordinate]int)

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

	for {
		var x IPType
		select {
		case x = <-computer.Output:
		case <-computer.Stopped:
			return grid
		}
		y := <-computer.Output
		tileId := <-computer.Output
		coord := Coordinate{
			X: int(x),
			Y: int(y),
		}
		grid[coord] = int(tileId)
	}
}

func main() {
	program := ParseProgram(adventutil.Parse(13)[0])
	grid := buildGrid(program)

	numBlocks := 0
	for _, tileId := range grid {
		if tileId == 2 {
			numBlocks++
		}
	}
	fmt.Printf("Part 1:%d\n", numBlocks)
}
