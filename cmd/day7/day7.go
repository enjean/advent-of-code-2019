package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"github.com/enjean/advent-of-code-2019/internal/intcode"
)

func CalculateThrusterSignal(phases [5]int, program []int) int {
	outputToThruster := make(chan int)

	var amplifiers [5]intcode.Computer
	for i := 0; i < 5; i++ {
		amplifiers[i] = intcode.CreateComputer(fmt.Sprintf("%d", i),
			map[int]func(intcode.Computer, []int, int) int{
				1: intcode.Add,
				2: intcode.Multiply,
				3: intcode.Save,
				4: intcode.PrintFunc,
				5: intcode.JumpIfTrue,
				6: intcode.JumpIfFalse,
				7: intcode.LessThan,
				8: intcode.Equals,
			})
	}
	for i := 0; i < 5; i++ {
		go func(ampIndex int) {
			amplifiers[ampIndex].Input <- phases[ampIndex]
			if ampIndex == 0 {
				amplifiers[ampIndex].Input <- 0
			}
			inputAmplifier := ampIndex - 1
			if ampIndex == 0 {
				inputAmplifier = 4
			}
			var output int
			for val := range amplifiers[inputAmplifier].Output {
				if inputAmplifier == 4 {
					output = val
				}
				select {
				case amplifiers[ampIndex].Input <- val:
				case <-amplifiers[ampIndex].Stopped:
				}
			}
			if inputAmplifier == 4 {
				outputToThruster <- output
			}
		}(i)
		go func(j int) {
			amplifiers[j].Run(program)
			//fmt.Printf("%d done\n", j)
		}(i)
	}

	return <-outputToThruster
}

func OptimalThrusterSignal(program []int, minPhase int) ([5]int, int) {
	phaserPermutations := generatePermutations(minPhase)

	var maxPhaserSetting [5]int
	maxThruster := 0
	for _, phaserPermutation := range phaserPermutations {
		thrusterSignal := CalculateThrusterSignal(phaserPermutation, program)
		//fmt.Printf("%v = %d\n", phaserPermutation, thrusterSignal)
		if thrusterSignal > maxThruster {
			maxThruster = thrusterSignal
			maxPhaserSetting = phaserPermutation
		}
	}
	return maxPhaserSetting, maxThruster
}

func generatePermutations(min int) [][5]int {
	var permutations [][5]int
	for a := min; a < min+5; a++ {
		for b := min; b < min+5; b++ {
			if a == b {
				continue
			}
			for c := min; c < min+5; c++ {
				if c == a || c == b {
					continue
				}
				for d := min; d < min+5; d++ {
					if d == a || d == b || d == c {
						continue
					}
					for e := min; e < min+5; e++ {
						if e == a || e == b || e == c || e == d {
							continue
						}
						permutations = append(permutations, [5]int{a, b, c, d, e})
					}
				}
			}
		}
	}
	return permutations
}

// Uses Heap's algorithm
// TODO not working
//func generatePermutations(k int, A [5]int) [][5]int {
//	fmt.Printf("generate(%d, %v)\n", k, A)
//	if k == 1 {
//		fmt.Printf("*%v*\n", A)
//		return [][5]int{A}
//	}
//	var permutations [][5]int
//	permutations = append(permutations, generatePermutations(k-1, A)...)
//	for i := 0; i < k-1; i++ {
//		if k%2 == 0 {
//			A[i], A[k-1] = A[k-1], A[i]
//		} else {
//			A[0], A[k-1] = A[k-1], A[0]
//		}
//		permutations = append(permutations, generatePermutations(k-1, A)...)
//	}
//	return permutations
//}

func main() {
	program := intcode.ParseProgram(adventutil.Parse(7)[0])

	_, part1 := OptimalThrusterSignal(program, 0)
	fmt.Printf("Part 1: %d \n", part1)

	_, part2 := OptimalThrusterSignal(program, 5)
	fmt.Printf("Part 2: %d \n", part2)
}
