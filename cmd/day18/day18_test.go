package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMinStepsToFindAllKeys(t *testing.T) {
	doMinStepsTest := func (input []string, expected int) {
		tunnelMap := parseMap(input)
		result := MinStepsToFindAllKeys(tunnelMap)
		if result != expected {
			t.Errorf("Expected %d got %d", expected, result)
		}
	}
	t.Run("Example 1", func(t *testing.T) {
		input := []string {
			"#########",
			"#b.A.@.a#",
			"#########",
		}
		doMinStepsTest(input, 8)
	})
	t.Run("Example 2", func(t *testing.T) {
		input := []string {
			"########################",
			"#f.D.E.e.C.b.A.@.a.B.c.#",
			"######################.#",
			"#d.....................#",
			"########################",
		}
		doMinStepsTest(input, 86)
	})
	t.Run("Example 3", func(t *testing.T) {
		input := []string {
			"########################",
			"#...............b.C.D.f#",
			"#.######################",
			"#.....@.a.B.c.d.A.e.F.g#",
			"########################",
		}
		doMinStepsTest(input, 132)
	})
	t.Run("Example 4", func(t *testing.T) {
		input := []string {
			"#################",
			"#i.G..c...e..H.p#",
			"########.########",
			"#j.A..b...f..D.o#",
			"########@########",
			"#k.E..a...g..B.n#",
			"########.########",
			"#l.F..d...h..C.m#",
			"#################",
		}
		doMinStepsTest(input, 136)
	})
	t.Run("Example 5", func(t *testing.T) {
		input := []string {
			"########################",
			"#@..............ac.GI.b#",
			"###d#e#f################",
			"###A#B#C################",
			"###g#h#i################",
			"########################",
		}
		start := time.Now()
		doMinStepsTest(input, 81)
		fmt.Printf("Took: %s\n", time.Since(start))
	})
}
