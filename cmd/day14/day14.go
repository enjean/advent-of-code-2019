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

func OreNeededForFuel(reactionStrings []string) int {
	reactions := parseReactionList(reactionStrings)
	oreNeeded := 0
	needed := []amountOfX{{1, "FUEL"}}
	leftovers := make(map[string]int)
	for len(needed) > 0 {
		fmt.Printf("%v %v %d\n", needed, leftovers, oreNeeded)
		processing := needed[0]
		needed = needed[1:]

		takenFromLeftovers := adventutil.Min(processing.amount, leftovers[processing.name])
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
				oreNeeded += ingredientAmountNeeded
				continue
			}

			//alreadyInQueue := false
			//for i, entry := range needed {
			//	if entry.name == ingredient.name {
			//		alreadyInQueue = true
			//		needed[i].amount += amountNeeded
			//	}
			//}
			//if alreadyInQueue {
			//	continue
			//}
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
	part1 := OreNeededForFuel(reactionStrings)
	fmt.Printf("Part 1: %d\n", part1)
}
