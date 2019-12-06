package main

import "testing"

func TestCountOrbits(t *testing.T) {
	input := []string {
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}

	result := CountOrbits(input)

	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
}

func TestOrbitalTransfersNeeded(t *testing.T) {
	input := []string {
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	result := OrbitalTransfersNeeded(input, "YOU", "SAN")

	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}
