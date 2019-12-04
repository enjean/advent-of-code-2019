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
