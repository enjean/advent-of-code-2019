package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"strconv"
)

type Layer [][]int

func ParseLayers(input string, width, height int) []Layer {
	var layers []Layer

	for len(input) > 0 {
		var layer Layer
		for r := 0; r < height; r++ {
			var row []int
			for c := 0; c < width; c++ {
				digitStr := input[0]
				input = input[1:]
				digit, _ := strconv.Atoi(string(digitStr))
				row = append(row, digit)
			}
			layer = append(layer, row)
		}
		layers = append(layers, layer)
	}

	return layers
}

func Part1(layers []Layer) int {
	var countsForlayerWithLeastZeros [10]int
	minZerosSeen := adventutil.MaxInt
	for _, layer := range layers {
		counts := layer.digitCounts()
		if counts[0] < minZerosSeen {
			minZerosSeen = counts[0]
			countsForlayerWithLeastZeros = counts
		}
	}
	return countsForlayerWithLeastZeros[1] * countsForlayerWithLeastZeros[2]
}

func Part2(layers []Layer, width, height int) {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			var character string
			for _, layer := range layers {
				if layer[r][c] == 0 {
					character = " "
					break
				}
				if layer[r][c] == 1 {
					character = "W"
					break
				}
			}
			fmt.Printf("%s", character)
		}
		fmt.Println()
	}

}

func (layer Layer) digitCounts() [10]int {
	var counts [10]int
	for _, row := range layer {
		for _, val := range row {
			counts[val]++
		}
	}
	return counts
}

func main() {
	layers := ParseLayers(adventutil.Parse(8)[0], 25, 6)
	part1 := Part1(layers)
	fmt.Printf("Part 1: %d\n", part1)
	Part2(layers, 25, 6)
}
