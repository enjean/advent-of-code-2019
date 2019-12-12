package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	. "github.com/enjean/advent-of-code-2019/internal/adventutil/coordinate"
	"math"
	"sort"
)

func ParseAsteroidLocations(lines []string) []Coordinate {
	var asteroidLocations []Coordinate
	for y, line := range lines {
		for x, character := range line {
			if character == '#' {
				asteroidLocations = append(asteroidLocations, Coordinate{x, y})
			}
		}
	}
	return asteroidLocations
}

func NumCanSee(source Coordinate, asteroids []Coordinate) int {
	numCanSee := 0
	for _, target := range asteroids {
		if source == target {
			continue
		}
		if canSee(source, target, asteroids) {
			numCanSee++
		}
	}
	return numCanSee
}

func canSee(source, target Coordinate, asteroids []Coordinate) bool {
	canSee := true
	for _, maybeInBetween := range asteroids {
		if source == maybeInBetween || target == maybeInBetween {
			continue
		}
		if isInBetween(source, target, maybeInBetween) {
			//fmt.Printf("%v can't see %v because %v is in between\n", source, target, maybeInBetween)
			canSee = false
			break
		}
	}
	return canSee
}

func isInBetween(point1, point2, candidate Coordinate) bool {
	// https://stackoverflow.com/questions/11907947/how-to-check-if-a-point-lies-on-a-line-between-2-other-points
	dxc := candidate.X - point1.X
	dyc := candidate.Y - point1.Y

	dxl := point2.X - point1.X
	dyl := point2.Y - point1.Y

	cross := dxc*dyl - dyc*dxl

	// If cross == 0 not on the same line
	if cross != 0 {
		return false
	}

	if adventutil.Abs(dxl) >= adventutil.Abs(dyl) {
		if dxl > 0 {
			return point1.X <= candidate.X && candidate.X <= point2.X
		}
		return point2.X <= candidate.X && candidate.X <= point1.X
	}
	if dyl > 0 {
		return point1.Y <= candidate.Y && candidate.Y <= point2.Y
	}
	return point2.Y <= candidate.Y && candidate.Y <= point1.Y
}

func FindMonitoringStation(regionMap []string) (Coordinate, int) {
	asteroids := ParseAsteroidLocations(regionMap)
	var canSeeMost Coordinate
	maxCanSee := 0
	for _, asteroid := range asteroids {
		canSee := NumCanSee(asteroid, asteroids)
		if canSee > maxCanSee {
			canSeeMost = asteroid
			maxCanSee = canSee
		}
	}
	return canSeeMost, maxCanSee
}

func AsteroidsDestroyed(asteroids []Coordinate, laser Coordinate) []Coordinate {
	for i, c := range asteroids {
		if c == laser {
			asteroids = append(asteroids[:i], asteroids[i+1:]...)
			break
		}
	}

	maxX := 0
	maxY := 0
	asteroidsMap := make(map[Coordinate]bool)
	for _, c := range asteroids {
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
		asteroidsMap[c] = true
	}

	sort.Slice(asteroids, func(i, j int) bool {
		dx1, dy1 := laser.DxDy(asteroids[i])
		dx2, dy2 := laser.DxDy(asteroids[j])
		q1 := quadrant(dx1, dy1)
		q2 := quadrant(dx2, dy2)

		if q1 < q2 {
			return true
		}
		if q1 > q2 {
			return false
		}
		slope1 := slope(dx1, dy1)
		slope2 := slope(dx2, dy2)
		if slope1 == slope2 {
			return laser.Distance(asteroids[i]) < laser.Distance(asteroids[j])
		}
		return slope1 < slope2
	})

	//for _, a := range asteroids {
	//	dx, dy := laser.DxDy(a)
	//	fmt.Printf("%s %d,%d %d %f \n", a, dx, dy, quadrant(dx, dy), slope(dx, dy))
	//}

	var destroyedOrder []Coordinate
	destroyedMap := make(map[Coordinate]bool)
	for len(destroyedMap) < len(asteroids) {
		var destroyedThisRound []Coordinate
		for _, asteroid := range asteroids {
			if destroyedMap[asteroid] {
				continue
			}
			isBlockedThisRound := false
			for _, d := range destroyedThisRound {
				if isInBetween(laser, asteroid, d) {
					isBlockedThisRound = true
					break
				}
			}
			if isBlockedThisRound {
				//fmt.Println("Would have destroyed but blocked " + asteroid.String())
				continue
			}
			//fmt.Printf("At %d Destroying %s\n", len(destroyedMap), asteroid)
			destroyedOrder = append(destroyedOrder, asteroid)
			destroyedMap[asteroid] = true
			destroyedThisRound = append(destroyedThisRound, asteroid)
			//printMap(asteroidsMap, destroyedMap, laser, maxX, maxY)
		}
	}

	return destroyedOrder
}

func printMap(asteroids, destroyed map[Coordinate]bool, laser Coordinate, maxX, maxY int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			var character string
			coord := Coordinate{x, y}
			if destroyed[coord] {
				character = "*"
			} else if asteroids[coord] {
				character = "#"
			} else if coord == laser {
				character = "X"
			} else {
				character = "."
			}
			fmt.Print(character)
		}
		fmt.Println()
	}
}

func quadrant(dx, dy int) int {
	if dx >= 0 && dy < 0 {
		return 1
	} else if dx > 0 && dy >= 0 {
		return 2
	} else if dx <= 0 && dy > 0 {
		return 3
	} else if dx < 0 && dy <= 0 {
		return 4
	} else {
		panic("Not assigned to quadrant")
	}
}

func slope(dx, dy int) float64 {
	if dx == 0 {
		return -math.MaxFloat32
	}
	return float64(dy) / float64(dx)
}

func main() {
	regionMap := adventutil.Parse(10)
	station, canSee := FindMonitoringStation(regionMap)
	fmt.Printf("Part 1: %s can see %d\n", station, canSee)

	destructionOrder := AsteroidsDestroyed(ParseAsteroidLocations(regionMap), station)
	twoHundredth := destructionOrder[199]
	part2 := 100 * twoHundredth.X + twoHundredth.Y
	fmt.Printf("Part 2: %d\n", part2)
}
