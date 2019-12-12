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
