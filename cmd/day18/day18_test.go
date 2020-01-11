package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMinStepsToFindAllKeys(t *testing.T) {
	doMinStepsTest := func(input []string, expected int) {
		tunnelMap := parseMap(input)
		result := MinStepsToFindAllKeys(tunnelMap)
		if result != expected {
			t.Errorf("Expected %d got %d", expected, result)
		}
	}
	t.Run("Example 1", func(t *testing.T) {
		input := []string{
			"#########",
			"#b.A.@.a#",
			"#########",
		}
		doMinStepsTest(input, 8)
	})
	t.Run("Example 2", func(t *testing.T) {
		input := []string{
			"########################",
			"#f.D.E.e.C.b.A.@.a.B.c.#",
			"######################.#",
			"#d.....................#",
			"########################",
		}
		doMinStepsTest(input, 86)
	})
	t.Run("Example 3", func(t *testing.T) {
		input := []string{
			"########################",
			"#...............b.C.D.f#",
			"#.######################",
			"#.....@.a.B.c.d.A.e.F.g#",
			"########################",
		}
		doMinStepsTest(input, 132)
	})
	t.Run("Example 4", func(t *testing.T) {
		input := []string{
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
		start := time.Now()
		doMinStepsTest(input, 136)
		fmt.Printf("Took: %s\n", time.Since(start))
	})
	t.Run("Example 5", func(t *testing.T) {
		input := []string{
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

	t.Run("a test", func(t *testing.T) {
		start := time.Now()
		for a := 0; a < 3000000000; a++ {
			for b := 0; b < 5; b++ {
				_ = a + b
			}
		}
		fmt.Printf("Took: %s\n", time.Since(start))

	})
}

func TestKeyAsIntExpr(t *testing.T) {
	tests := []struct {
		key      rune
		expected int
	}{
		{'a', 1},
		{'b', 2},
		{'c', 4},
		{'d', 8},
	}
	for _, test := range tests {
		result := keyAsIntExpr(test.key)
		if result != test.expected {
			t.Errorf("key(%s) expected %d got %d", string(test.key), test.expected, result)
		}
	}
}

func TestHasRequiredKeys(t *testing.T) {
	toKeyExpr := func (s string) int {
		result := 0
		for _, r := range s {
			result += keyAsIntExpr(r)
		}
		return result
	}
	tests := []struct {
		keysPossessed string
		keysNeeded string
		expected bool
	}{
		{"", "", true},
		{"", "a", false},
		{"a", "a", true},
		{"a", "b", false},
		{"ab", "de", false},
		{"ab", "ba", true},
		{"abcde", "ce", true},
	}
	for _, test := range tests {
		result := hasRequiredKeys(toKeyExpr(test.keysPossessed), toKeyExpr(test.keysNeeded))
		if result != test.expected {
			t.Errorf("hasRequiredKeys(%s, %s) expected %v got %v", test.keysPossessed, test.keysNeeded, test.expected, result)
		}
	}
}
