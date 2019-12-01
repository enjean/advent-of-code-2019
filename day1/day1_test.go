package main

import "testing"

func TestFuelNeeded(t *testing.T) {
	tests := []struct {
		input int
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	//For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
	//For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
	//For a mass of 1969, the fuel required is 654.
	//For a mass of 100756, the fuel required is 33583.
	for _, test := range tests {
		result := FuelNeeded(test.input)
		if result != test.expected {
			t.Errorf("FuelNeeded(%d) expected %d, got %d", test.input, test.expected, result)
		}
	}
}
