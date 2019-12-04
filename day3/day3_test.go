package main

import "testing"

func TestFindNearestIntersection(t *testing.T) {
	tests := []struct {
		wires    []Wire
		expected int
	}{
		{
			[]Wire{
				{{R, 8}, {U, 5}, {L, 5}, {D, 3},},
				{{U, 7}, {R, 6}, {D, 4}, {L, 4},}, //U7,R6,D4,L4
			},
			6,
		},
		{
			[]Wire{
				//R75,D30,R83,U83,L12,D49,R71,U7,L72
				//U62,R66,U55,R34,D71,R55,D58,R83
				{{R, 75}, {D, 30}, {R, 83}, {U, 83}, {L, 12}, {D, 49}, {R, 71}, {U, 7}, {L, 72},},
				{{U, 62}, {R, 66}, {U, 55}, {R, 34}, {D, 71}, {R, 55}, {D, 58}, {R, 83},},
			},
			159,
		},
		{
			[]Wire{
				{{R, 98}, {U, 47}, {R, 26}, {D, 63}, {R, 33}, {U, 87}, {L, 62}, {D, 20}, {R, 33}, {U, 53}, {R, 51},},
				{{U, 98}, {R, 91}, {D, 20}, {R, 16}, {D, 67}, {R, 40}, {U, 7}, {R, 15}, {U, 6}, {R, 7}},
			},
			135,
		},
	}
	for _, test := range tests {
		result := FindNearestIntersection(test.wires)
		if result != test.expected {
			t.Errorf("Nearest intersection %v expected %d got %d", test.wires, test.expected, result)
		}
	}
}

func TestFindFirstIntersection(t *testing.T) {
	tests := []struct {
		wire1, wire2 Wire
		expected     int
	}{
		{
			Wire{{R, 8}, {U, 5}, {L, 5}, {D, 3},},
			Wire{{U, 7}, {R, 6}, {D, 4}, {L, 4},}, //U7,R6,D4,L4
			30,
		},
		{
			Wire{{R, 75}, {D, 30}, {R, 83}, {U, 83}, {L, 12}, {D, 49}, {R, 71}, {U, 7}, {L, 72},},
			Wire{{U, 62}, {R, 66}, {U, 55}, {R, 34}, {D, 71}, {R, 55}, {D, 58}, {R, 83},},
			610,
		},
		{
			Wire{{R, 98}, {U, 47}, {R, 26}, {D, 63}, {R, 33}, {U, 87}, {L, 62}, {D, 20}, {R, 33}, {U, 53}, {R, 51},},
			Wire{{U, 98}, {R, 91}, {D, 20}, {R, 16}, {D, 67}, {R, 40}, {U, 7}, {R, 15}, {U, 6}, {R, 7}},
			410,
		},
	}
	for _, test := range tests {
		result := FindFirstIntersection(test.wire1, test.wire2)
		if result != test.expected {
			t.Errorf("Nearest intersection %v %v expected %d got %d", test.wire1, test.wire2, test.expected, result)
		}
	}
}
