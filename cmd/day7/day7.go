package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"github.com/enjean/advent-of-code-2019/internal/intcode"
)

func CalculateThrusterSignal(phases [5]int, program []int) int {
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

		go func(ampIndex int) {
			amplifiers[ampIndex].Input <- phases[ampIndex]
			if ampIndex == 0 {
				amplifiers[ampIndex].Input <- 0
			}
			if ampIndex != 0 {
				for val := range amplifiers[ampIndex-1].Output {
					amplifiers[ampIndex].Input <- val
				}
			}
		}(i)

		go func(j int) {
			amplifiers[j].Run(program)
			//fmt.Printf("%d done\n", j)
		}(i)
	}

	return <-amplifiers[4].Output
}

func OptimalThrusterSignal(program []int) ([5]int, int) {
	phaserPossibilities := [5]int{0, 1, 2, 3, 4}
	phaserPermutations := generatePermutations(5, phaserPossibilities)

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

func generatePermutations(k int, A [5]int) [][5]int {
	var permutations [][5]int
	for a := 0; a<5; a++ {
		for b := 0; b<5;b++ {
			if a==b {
				continue
			}
			for c:=0;c<5;c++ {
				if c==a || c==b {
					continue
				}
				for d:=0; d<5;d++{
					if d==a||d==b||d==c {
						continue
					}
					for e:=0;e<5;e++ {
						if e==a||e==b||e==c||e==d {
							continue
						}
						permutations = append(permutations, [5]int{a,b,c,d,e})
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

	_, part1 := OptimalThrusterSignal(program)
	fmt.Printf("Part 1: %d \n", part1)
}
