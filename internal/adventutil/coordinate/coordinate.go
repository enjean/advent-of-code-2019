package coordinate

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"math"
)

type Coordinate struct {
	X, Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func (c Coordinate) DxDy(o Coordinate) (int, int) {
	return o.X - c.X, o.Y - c.Y
}

func (c Coordinate) Distance(o Coordinate) float64 {
	return math.Sqrt(math.Pow(float64(o.X-c.X), 2) + math.Pow(float64(o.Y-c.Y), 2))
}

func (c Coordinate) ManhattanDistance(o Coordinate) int {
	return adventutil.Abs(c.X-o.X) + adventutil.Abs(c.Y-o.Y)
}

func (c Coordinate) Neighbors() []Coordinate {
	return []Coordinate{
		{c.X, c.Y - 1},
		{c.X + 1, c.Y},
		{c.X, c.Y + 1},
		{c.X - 1, c.Y},
	}
}

func PrintIntCoordinateMap(coordMap map[Coordinate]int, valToString func(int) string) {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	for c := range coordMap {
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
			fmt.Print(valToString(coordMap[Coordinate{X: x, Y: y}]))
		}
		fmt.Println()
	}
}
