package main

import "testing"

func TestIsInBetween(t *testing.T) {
	if !isInBetween(Coordinate{5,8},Coordinate{0,8},Coordinate{1,8}) {
		t.Errorf("Should be in between")
	}
}

func TestCanSee(t *testing.T) {
	input := []string{
		"......#.#.",
		"#..#.#....",
		"..#######.",
		".#.#.###..",
		".#..#.....",
		"..#....#.#",
		"#..#....#.",
		".##.#..###",
		"##...#..#.",
		".#....####",
	}
	asteroids := ParseAsteroidLocations(input)
	source := Coordinate{5, 8}
	target := Coordinate{0,8}
	if canSee(source, target, asteroids) {
		t.Errorf("Should not be able to see")
	}
}

func TestNumCanSee(t *testing.T) {
	tests := []struct {
		input        []string
		coord        Coordinate
		expectedSeen int
	}{
		{
			[]string{
				".#..#",
				".....",
				"#####",
				"....#",
				"...##",
			},
			Coordinate{3, 4},
			8,
		},
		{
			[]string{
				"......#.#.",
				"#..#.#....",
				"..#######.",
				".#.#.###..",
				".#..#.....",
				"..#....#.#",
				"#..#....#.",
				".##.#..###",
				"##...#..#.",
				".#....####",
			},
			Coordinate{5, 8},
			33,
		},
		{
			[]string{
				"#.#...#.#.",
				".###....#.",
				".#....#...",
				"##.#.#.#.#",
				"....#.#.#.",
				".##..###.#",
				"..#...##..",
				"..##....##",
				"......#...",
				".####.###.",
			},
			Coordinate{1, 2},
			35,
		},
		{
			[]string{
				".#..#..###",
				"####.###.#",
				"....###.#.",
				"..###.##.#",
				"##.##.#.#.",
				"....###..#",
				"..#.#..#.#",
				"#..#.#.###",
				".##...##.#",
				".....#.#..",
			},
			Coordinate{6, 3},
			41,
		},
		{
			[]string{
				".#..##.###...#######",
				"##.############..##.",
				".#.######.########.#",
				".###.#######.####.#.",
				"#####.##.#.##.###.##",
				"..#####..#.#########",
				"####################",
				"#.####....###.#.#.##",
				"##.#################",
				"#####.##.###..####..",
				"..######..##.#######",
				"####.##.####...##..#",
				".#####..#.######.###",
				"##...#.##########...",
				"#.##########.#######",
				".####.#.###.###.#.##",
				"....##.##.###..#####",
				".#.#.###########.###",
				"#.#.#.#####.####.###",
				"###.##.####.##.#..##",
			},
			Coordinate{11, 13},
			210,
		},
	}
	for _, test := range tests {
		asteroidLocations := ParseAsteroidLocations(test.input)
		seen := NumCanSee(test.coord, asteroidLocations)
		if seen != test.expectedSeen {
			t.Errorf("Got incorrect seen %d, expected %d", seen, test.expectedSeen)
		}
	}
}

func TestFindMonitoringStation(t *testing.T) {
	tests := []struct {
		input         []string
		expectedCoord Coordinate
		expectedSeen  int
	}{
		{
			[]string{
				".#..#",
				".....",
				"#####",
				"....#",
				"...##",
			},
			Coordinate{3, 4},
			8,
		},
		{
			[]string{
				"......#.#.",
				"#..#.#....",
				"..#######.",
				".#.#.###..",
				".#..#.....",
				"..#....#.#",
				"#..#....#.",
				".##.#..###",
				"##...#..#.",
				".#....####",
			},
			Coordinate{5, 8},
			33,
		},
		{
			[]string{
				"#.#...#.#.",
				".###....#.",
				".#....#...",
				"##.#.#.#.#",
				"....#.#.#.",
				".##..###.#",
				"..#...##..",
				"..##....##",
				"......#...",
				".####.###.",
			},
			Coordinate{1, 2},
			35,
		},
		{
			[]string{
				".#..#..###",
				"####.###.#",
				"....###.#.",
				"..###.##.#",
				"##.##.#.#.",
				"....###..#",
				"..#.#..#.#",
				"#..#.#.###",
				".##...##.#",
				".....#.#..",
			},
			Coordinate{6, 3},
			41,
		},
		{
			[]string{
				".#..##.###...#######",
				"##.############..##.",
				".#.######.########.#",
				".###.#######.####.#.",
				"#####.##.#.##.###.##",
				"..#####..#.#########",
				"####################",
				"#.####....###.#.#.##",
				"##.#################",
				"#####.##.###..####..",
				"..######..##.#######",
				"####.##.####...##..#",
				".#####..#.######.###",
				"##...#.##########...",
				"#.##########.#######",
				".####.#.###.###.#.##",
				"....##.##.###..#####",
				".#.#.###########.###",
				"#.#.#.#####.####.###",
				"###.##.####.##.#..##",
			},
			Coordinate{11, 13},
			210,
		},
	}
	for _, test := range tests {
		coord, seen := FindMonitoringStation(test.input)
		if coord != test.expectedCoord {
			t.Errorf("Got incorrect coord %v expected %v", coord, test.expectedCoord)
		}
		if seen != test.expectedSeen {
			t.Errorf("Got incorrect seen %d, expected %d", seen, test.expectedSeen)
		}
	}
}
