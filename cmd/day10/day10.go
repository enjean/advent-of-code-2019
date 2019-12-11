package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
)

type Coordinate struct {
	X, Y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

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

func main() {
	regionMap := adventutil.Parse(10)
	_, canSee := FindMonitoringStation(regionMap)
	fmt.Printf("Part 1:%d\n", canSee)
}
