package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"regexp"
	"strconv"
)

type ThreeDCoord struct {
	x, y, z int
}

func (c ThreeDCoord) add(o ThreeDCoord) ThreeDCoord {
	return ThreeDCoord{
		x: c.x + o.x,
		y: c.y + o.y,
		z: c.z + o.z,
	}
}

type Moon struct {
	position, velocity ThreeDCoord
}

func (m Moon) potentialEnergy() int {
	return adventutil.Abs(m.position.x) +
		adventutil.Abs(m.position.y) +
		adventutil.Abs(m.position.z)
}

func (m Moon) kineticEnergy() int {
	return adventutil.Abs(m.velocity.x) +
		adventutil.Abs(m.velocity.y) +
		adventutil.Abs(m.velocity.z)
}

func (m Moon) totalEnergy() int {
	return m.potentialEnergy() * m.kineticEnergy()
}

type MoonSystem struct {
	moons []*Moon
}

func (ms MoonSystem) TotalEnergy() int {
	total := 0
	for _, moon := range ms.moons {
		total += moon.totalEnergy()
	}
	return total
}

// <x=-1, y=0, z=2>
var moonRegex = regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)

func ParseMoonSystem(lines []string) *MoonSystem {
	var moons []*Moon
	for _, line := range lines {
		matches := moonRegex.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		z, _ := strconv.Atoi(matches[3])
		moons = append(moons, &Moon{position: ThreeDCoord{x: x, y: y, z: z,}})
	}
	return &MoonSystem{moons: moons}
}

func (ms *MoonSystem) ApplyStep() {
	for i, moon := range ms.moons {
		dVx, dVy, dVz := 0, 0, 0
		for j, otherMoon := range ms.moons {
			if i == j {
				continue
			}
			if moon.position.x < otherMoon.position.x {
				dVx++
			} else if moon.position.x > otherMoon.position.x {
				dVx--
			}
			if moon.position.y < otherMoon.position.y {
				dVy++
			} else if moon.position.y > otherMoon.position.y {
				dVy--
			}
			if moon.position.z < otherMoon.position.z {
				dVz++
			} else if moon.position.z > otherMoon.position.z {
				dVz--
			}
		}
		moon.velocity = moon.velocity.add(ThreeDCoord{x: dVx, y: dVy, z: dVz})
	}

	for _, moon := range ms.moons {
		moon.position = moon.position.add(moon.velocity)
	}
}

func Simulate(ms *MoonSystem, steps int) {
	for step := 0; step < steps; step++ {
		ms.ApplyStep()
	}
}

func main() {
	moonSystem := ParseMoonSystem(adventutil.Parse(12))
	Simulate(moonSystem, 1000)
	fmt.Printf("Part 1: %d\n", moonSystem.TotalEnergy())
}
