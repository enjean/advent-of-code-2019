package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	"math"
	"strconv"
	"strings"
)

type amountOfX struct {
	amount int
	name   string
}

type equation struct {
	ingredients []amountOfX
	produces    int
}

func parseReactionList(reactionStrings []string) map[string]equation {
	reactions := make(map[string]equation)
	for _, line := range reactionStrings {
		lhsRhs := strings.Split(line, " => ")
		rhs := parseTerm(lhsRhs[1])

		var ingredients []amountOfX
		for _, lhsTerm := range strings.Split(lhsRhs[0], ", ") {
			ingredients = append(ingredients, parseTerm(lhsTerm))
		}
		reactions[rhs.name] = equation{
			ingredients: ingredients,
			produces:    rhs.amount,
		}
	}
	return reactions
}

func OreNeededForFuel(reactions map[string]equation, fuel int) int {
	oreNeeded := 0
	needed := []amountOfX{{fuel, "FUEL"}}
	leftovers := make(map[string]int)
	for len(needed) > 0 {
		//fmt.Printf("%v %v %d\n", needed, leftovers, oreNeeded)
		processing := needed[0]
		needed = needed[1:]

		takenFromLeftovers := adventutil.Min(processing.amount, leftovers[processing.name])
		//if takenFromLeftovers > 0 {
		//	fmt.Println("Took from leftovers " + processing.name)
		//}
		leftovers[processing.name] -= takenFromLeftovers
		amountNeeded := processing.amount - takenFromLeftovers
		if amountNeeded == 0 {
			continue
		}

		equation := reactions[processing.name]

		multiplier := int(math.Ceil(float64(amountNeeded) / float64(equation.produces)))

		if amountNeeded < equation.produces*multiplier {
			leftovers[processing.name] = equation.produces*multiplier - amountNeeded
		}
		for _, ingredient := range equation.ingredients {
			ingredientAmountNeeded := ingredient.amount * multiplier
			if ingredient.name == "ORE" {
				//fmt.Printf("Need %d ore to make %d %s\n", ingredientAmountNeeded, amountNeeded, processing.name)
				oreNeeded += ingredientAmountNeeded
				continue
			}
			needed = append(needed, amountOfX{
				amount: ingredientAmountNeeded,
				name:   ingredient.name,
			})

		}
	}
	return oreNeeded
}

func parseTerm(term string) amountOfX {
	parts := strings.Split(term, " ")
	amount, _ := strconv.Atoi(parts[0])
	return amountOfX{
		amount: amount,
		name:   parts[1],
	}
}

func main() {
	reactionStrings := adventutil.Parse(14)
	equations := parseReactionList(reactionStrings)
	part1 := OreNeededForFuel(equations, 1)
	fmt.Printf("Part 1: %d\n", part1)

	//example := parseReactionList([]string{
	//	"157 ORE => 5 NZVS",
	//	"165 ORE => 6 DCFZ",
	//	"44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL",
	//	"12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ",
	//	"179 ORE => 7 PSHF",
	//	"177 ORE => 5 HKGWZ",
	//	"7 DCFZ, 7 PSHF => 2 XJWVT",
	//	"165 ORE => 2 GPVTF",
	//	"3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT",
	//})
	//for i:=1; i<=50; i++ {
	//	fmt.Printf("%d,%d\n", i, OreNeededForFuel(example, i))
	//}

	guess := 1000000000000 / part1
	needed := OreNeededForFuel(equations, guess)
	//fmt.Printf("%d produces %d\n", needed, guess)
	//for needed < 1000000000000 {
	//	guess += 100
	//	needed = OreNeededForFuel(equations, guess)
	//	fmt.Printf("%d produces %d\n", needed, guess)
	//}

	guess = 8193591
	for needed < 1000000000000 {
		guess += 1
		needed = OreNeededForFuel(equations, guess)
		fmt.Printf("%d produces %d\n", needed, guess)
	}
}
