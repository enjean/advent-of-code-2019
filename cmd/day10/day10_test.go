package main

import (
	. "github.com/enjean/advent-of-code-2019/internal/adventutil/coordinate"
	"testing"
)

func TestIsInBetween(t *testing.T) {
	if !isInBetween(Coordinate{5, 8}, Coordinate{0, 8}, Coordinate{1, 8}) {
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
	target := Coordinate{0, 8}
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

func TestAsteroidsDestoyed(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		asteroids := ParseAsteroidLocations([]string{
			".#....#####...#..",
			"##...##.#####..##",
			"##...#...#.#####.",
			"..#.....X...###..",
			"..#.#.....#....##",
		})
		station := Coordinate{8, 3}
		result := AsteroidsDestroyed(asteroids, station)
		//		.#....###24...#..
		//		##...##.13#67..9#
		//		##...#...5.8####.
		//		..#.....X...###..
		//		..#.#.....#....##

		//		.#....###.....#..
		//		##...##...#.....#
		//		##...#......1234.
		//		..#.....X...5##..
		//		..#.9.....8....76

		//.8....###.....#..
		//56...9#...#.....#
		//34...7...........
		//..2.....X....##..
		//..1..............

		//......234.....6..
		//......1...5.....7
		//.................
		//........X....89..
		//.................
		expected := []Coordinate{
			{8, 1}, {9, 0}, {9, 1}, {10, 0}, {9, 2}, {11, 1}, {12, 1}, {11, 2}, {15, 1},
			{12, 2}, {13, 2}, {14, 2}, {15, 2}, {12, 3}, {16, 4}, {15, 4}, {10, 4}, {4, 4},
			{2, 4}, {2, 3}, {0, 2}, {1, 2}, {0, 1}, {1, 1}, {5, 2}, {1, 0}, {5, 1},
			{6, 1}, {6, 0}, {7, 0}, {8, 0}, {10, 1}, {14, 0}, {16, 1}, {13, 3}, {14, 3},
		}
		if len(result) != len(expected) {
			t.Errorf("Didn't destroy correct number of asteroids expected %d got %d", len(expected), len(result))
		}
		for index, coord := range expected {
			if result[index] != coord {
				t.Errorf("Asteroid %d expected %s was %s", index, coord, result[index])
			}
		}
	})
	t.Run("Bigger", func(t *testing.T) {
		asteroids := ParseAsteroidLocations([]string{
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
		})
		station := Coordinate{11, 13}
		result := AsteroidsDestroyed(asteroids, station)

		//The 1st asteroid to be vaporized is at 11,12.
		//The 2nd asteroid to be vaporized is at 12,1.
		//The 3rd asteroid to be vaporized is at 12,2.
		//The 10th asteroid to be vaporized is at 12,8.
		//The 20th asteroid to be vaporized is at 16,0.
		//The 50th asteroid to be vaporized is at 16,9.
		//The 100th asteroid to be vaporized is at 10,16.
		//The 199th asteroid to be vaporized is at 9,6.
		//The 200th asteroid to be vaporized is at 8,2.
		//The 201st asteroid to be vaporized is at 10,9.
		//The 299th and final asteroid to be vaporized is at 11,1
		if len(result) != 299 {
			t.Errorf("Length wrong")
		}
		expected := map[int]Coordinate{
			0:   {11, 12},
			1:   {12, 1},
			2:   {12, 2},
			9:   {12, 8},
			19:  {16, 0},
			49:  {16, 9},
			99:  {10, 16},
			198: {9, 6},
			199: {8, 2},
			200: {10, 9},
			298: {11, 1},
		}
		for index, coord := range expected {
			if result[index] != coord {
				t.Errorf("Asteroid %d expected %s was %s", index, coord, result[index])
			}
		}

	})
	t.Run("Edge case", func(t *testing.T) {
		asteroids := ParseAsteroidLocations([]string{
			"..#..",
			"..#..",
			"##.##",
			"..#..",
			"..#..",
		})
		station := Coordinate{2, 2}
		result := AsteroidsDestroyed(asteroids, station)
		expected := []Coordinate{
			{2,1},{3,2},{2,3},{1,2},
			{2,0},{4,2},{2,4},{0,2},
		}
		if len(result) != len(expected) {
			t.Errorf("Didn't destroy correct number of asteroids expected %d got %d", len(expected), len(result))
		}
		for index, coord := range expected {
			if result[index] != coord {
				t.Errorf("Asteroid %d expected %s was %s", index, coord, result[index])
			}
		}
	})
	t.Run("q3 test", func(t *testing.T) {
		asteroids := ParseAsteroidLocations([]string{
			"...",
			"...",
			"##.",
		})
		station := Coordinate{1, 1}
		result := AsteroidsDestroyed(asteroids, station)
		expected := []Coordinate{
			{1,2},{0,2},
		}
		if len(result) != len(expected) {
			t.Errorf("Didn't destroy correct number of asteroids expected %d got %d", len(expected), len(result))
		}
		for index, coord := range expected {
			if result[index] != coord {
				t.Errorf("Asteroid %d expected %s was %s", index, coord, result[index])
			}
		}
	})
}
