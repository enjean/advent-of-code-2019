package main

import (
	"strconv"
	"testing"
)

func TestApplyPhase(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		val := parseSignal("12345678")
		expectedOutputs := []string {
			"48226158",
			"34040438",
			"03415518",
			"01029498",
		}
		for n, expected := range expectedOutputs {
			val = ApplyPhase(val)
			if asString(val) != expected {
				t.Errorf("After %d phase expected %s got %s", n, expected, asString(val))
			}
		}
	})
	t.Run("100 Run Tests", func(t *testing.T) {
		tests := []struct {
			input string
			expected string
		}{
			{"80871224585914546619083218645595","24176176"},
			{"19617804207202209144916044189917", "73745418"},
			{"69317163492948606335995924319873", "52432133"},
		}
		for _, test := range tests {
			val := parseSignal(test.input)
			for i := 0; i < 100; i++ {
				val = ApplyPhase(val)
			}
			first8 := asString(val)[:8]
			if first8 != test.expected {
				t.Errorf("Expected %s got %s", test.expected, first8)
			}
		}
	})
}

func asString(intArray []int) string {
	var result string
	for _, val := range intArray {
		result += strconv.Itoa(val)
	}
	return result
}