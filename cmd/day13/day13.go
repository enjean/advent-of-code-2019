package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	. "github.com/enjean/advent-of-code-2019/internal/adventutil/coordinate"
	. "github.com/enjean/advent-of-code-2019/internal/intcode"
)

func buildGrid(program Program) map[Coordinate]int {
	grid := make(map[Coordinate]int)

	computer := CreateCompleteComputer("Game")

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

func playGame(program Program) int {
	var score int
	grid := make(map[Coordinate]int)

	computer := CreateCompleteComputer("Game")

	program[0] = 2 // insert two quarters
	go func() { computer.Run(program) }()

	var paddleLocation Coordinate
	var ballLocation Coordinate
	walls := make(map[Coordinate]bool)
	joystickInput := 0

	for {

		if ballLocation.X > paddleLocation.X {
			joystickInput = 1
		} else if ballLocation.X < paddleLocation.X {
			joystickInput = -1
		} else {
			joystickInput = 0
		}

		var x IPType
		select {
		case <-computer.Stopped:
			return score
		case computer.Input <- IPType(joystickInput):
			//fmt.Printf("Sent input %d\n", joystickInput)
		case x = <-computer.Output:
			y := <-computer.Output
			tileId := <-computer.Output

			if x == -1 && y == 0 {
				score = int(tileId)
				continue
			}

			coord := Coordinate{
				X: int(x),
				Y: int(y),
			}
			//fmt.Printf("%s = %d\n", coord, tileId)
			grid[coord] = int(tileId)
			//if tileId != 0 {
			//	printGame(grid)
			//	fmt.Println()
			//}
			if tileId == 1 {
				walls[coord] = true
			}
			if tileId == 3 {
				paddleLocation = coord
			}
			if tileId == 4 {
				//if coord.Y > ballLocation.Y && ballLocation.Y != -1 {
				//	paddleTargetX = calculatePaddleTargetX(ballLocation, coord, paddleLocation.Y, walls)
				//}
				ballLocation = coord
			}
		}
	}
}

func calculatePaddleTargetX(lastBallLocation, currentBallLocation Coordinate, paddleY int, walls map[Coordinate]bool) int {
	dx := currentBallLocation.X - lastBallLocation.X
	x := currentBallLocation.X
	for y := currentBallLocation.Y; y < paddleY-1; y++ {
		if walls[Coordinate{X: x + 1, Y: y}] || walls[Coordinate{X: x - 1, Y: y}] {
			dx *= -1
		}
		x += dx
	}
	return x
}

func printGame(grid map[Coordinate]int) {
	PrintIntCoordinateMap(grid, func(i int) string {
		//0 is an empty tile. No game object appears in this tile.
		//1 is a wall tile. Walls are indestructible barriers.
		//2 is a block tile. Blocks can be broken by the ball.
		//3 is a horizontal paddle tile. The paddle is indestructible.
		//4 is a ball tile. The ball moves diagonally and bounces off objects.
		switch i {
		case 1:
			return "*"
		case 2:
			return "#"
		case 3:
			return "_"
		case 4:
			return "O"
		default:
			return " "
		}
	})
}

func main() {
	program := ParseProgram(adventutil.Parse(13)[0])
	//grid := buildGrid(program)

	//numBlocks := 0
	//for _, tileId := range grid {
	//	if tileId == 2 {
	//		numBlocks++
	//	}
	//}
	//fmt.Printf("Part 1:%d\n", numBlocks)
	score := playGame(program)
	fmt.Printf("Part 2: %d\n", score)
}
