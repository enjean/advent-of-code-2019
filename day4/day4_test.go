package main

import "testing"

func TestValidate(t *testing.T) {
	//111111 meets these criteria (double 11, never decreases).
	//223450 does not meet these criteria (decreasing pair of digits 50).
	//123789 does not meet these criteria (no double).
	tests := []struct{
		password string
		valid bool
	}{
		{"111111", true},
		{"223450", false},
		{"123789", false},
	}
	for _, test := range tests {
		if Validate(test.password) != test.valid {
			t.Errorf("Failed check %s", test.password)
		}
	}
}

func TestValidate2(t *testing.T) {
	//112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
	//123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
	//111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
	tests := []struct{
		password string
		valid bool
	}{
		{"112233", true},
		{"123444", false},
		{"111122", true},
	}
	for _, test := range tests {
		if Validate2(test.password) != test.valid {
			t.Errorf("Failed check %s", test.password)
		}
	}
}