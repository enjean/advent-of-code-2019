package main

import "testing"

func TestParseLayers(t *testing.T) {
	input := "123456789012"
	result := ParseLayers(input, 3, 2)
	expected := []Layer{
		{
			{1, 2, 3},
			{4, 5, 6},
		},
		{
			{7, 8, 9},
			{0, 1, 2},
		},
	}
	if len(result) != len(expected) {
		t.Errorf("Expected %d layers, was %d", len(expected), len(result))
	}
	for l, expectedLayer := range expected {
		resultLayer := result[l]
		if len(resultLayer) != len(expectedLayer) {
			t.Errorf("Expected %d rows, was %d", len(expectedLayer), len(resultLayer))
		}
		for r, expectedRow := range expectedLayer {
			resultRow := resultLayer[r]
			if len(resultRow) != len(expectedRow) {
				t.Errorf("Expected %d columns, was %d", len(expectedRow), len(resultRow))
			}
			for c, val := range expectedRow {
				if resultRow[c] != val {
					t.Errorf("Layer %d[%d][%d] expected %d, got %d", l, r, c, val, resultRow[c])
				}
			}
		}
	}
}
