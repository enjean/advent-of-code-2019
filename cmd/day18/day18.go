package main

import (
	"fmt"
	"github.com/enjean/advent-of-code-2019/internal/adventutil"
	. "github.com/enjean/advent-of-code-2019/internal/adventutil/coordinate"
	"math"
	"strings"
	"unicode"
)

func parseMap(input []string) map[Coordinate]rune {
	tunnelMap := make(map[Coordinate]rune)

	for y, line := range input {
		for x, val := range line {
			coord := Coordinate{
				X: x,
				Y: y,
			}
			tunnelMap[coord] = val
		}
	}

	return tunnelMap
}

type searchState struct {
	key       int
	keysFound int
	steps     int
	index     int
}

type PriorityQueue []*searchState

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].steps < pq[j].steps
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*searchState)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type adjacencyEntry struct {
	steps         int
	keysInBetween string
}

type adjacencyMatrix map[rune]map[rune]adjacencyEntry

func (am adjacencyMatrix) distanceBetween(source, dest rune, keysPossessed string) int {
	entry := am[source][dest]
	if hasRequiredKeys(keysPossessed, entry.keysInBetween) {
		return entry.steps
	}
	return adventutil.MaxInt
}

func MinStepsToFindAllKeys(tunnelMap map[Coordinate]rune) int {

	var keys []rune
	adjacencyMatrix := make(adjacencyMatrix)
	for coord, val := range tunnelMap {
		if val == '@' {
			adjacencyMatrix[0] = buildAdjacencyEntries(coord, tunnelMap)

		}
		if unicode.IsLower(val) {
			keys = append(keys, val)
			adjacencyMatrix[val] = buildAdjacencyEntries(coord, tunnelMap)
		}
	}

	for key, entries := range adjacencyMatrix {
		fmt.Println(string(key) + ":")
		for target, entry := range entries {
			fmt.Printf("  %s %d %s\n", string(target), entry.steps, entry.keysInBetween)
		}
	}

	lastCalculations := make(map[rune]map[string]int)
	for _, key := range keys {
		lastCalculations[key] = make(map[string]int)
		distanceBetween := adjacencyMatrix.distanceBetween(0, key, "")
		lastCalculations[key][""] = distanceBetween
		fmt.Printf("[%s, 0] = %d\n", string(key), distanceBetween)
	}

	for setSize := 1; setSize < len(keys); setSize++ {
		temp := make(map[rune]map[string]int)
		for _, key := range keys {
			temp[key] = make(map[string]int)
		}

		sets := setsOfSize(keys, setSize)
		for _, set := range sets {
			for _, key := range keys {
				if strings.ContainsRune(set, key) {
					continue
				}
				min := adventutil.MaxInt
				for _, parent := range set {
					parentToKey := adjacencyMatrix.distanceBetween(parent, key, set)
					if parentToKey == adventutil.MaxInt {
						continue
					}
					restOfSet := strings.Replace(set, string(parent), "", 1)
					restOfSetToParent := lastCalculations[parent][restOfSet]
					if restOfSetToParent == adventutil.MaxInt {
						continue
					}
					distance := parentToKey + restOfSetToParent
					if distance < min {
						min = distance
					}
				}
				//fmt.Printf("[%s, {%s}] = %d\n", string(key), set, min)
				temp[key][set] = min
			}
		}
		lastCalculations = temp
	}

	finalMin := adventutil.MaxInt
	for _, valsForKey := range lastCalculations {
		for _, valsForSet := range valsForKey {
			if valsForSet < finalMin {
				finalMin = valsForSet
			}
		}
	}

	return finalMin
}

func setsOfSize(vals []rune, size int) []string {
	// A temporary array to store all combination one by one
	setSoFar := make([]rune, size)
	var setsFound []string

	// Print all combination using temprary
	// array 'data[]'
	setFindUtil(vals, size, setSoFar, 0, 0, &setsFound)
	return setsFound
}

func setFindUtil(vals []rune, setSize int, setSoFar []rune, index, i int, setsFound *[]string) {
	if index == setSize {
		*setsFound = append(*setsFound, string(setSoFar))
		return
	}

	// When no more elements are there to put in data[]
	if i >= len(vals) {
		return
	}

	// current is included, put next at next location
	setSoFar[index] = vals[i]
	setFindUtil(vals, setSize, setSoFar, index+1, i+1, setsFound)

	// current is excluded, replace it with
	// next (Note that i+1 is passed, but
	// index is not changed)
	setFindUtil(vals, setSize, setSoFar, index, i+1, setsFound)
}

func buildAdjacencyEntries(start Coordinate, tunnelMap map[Coordinate]rune) map[rune]adjacencyEntry {
	adjacencyEntries := make(map[rune]adjacencyEntry)

	type pathState struct {
		coordinate    Coordinate
		steps         int
		keysInBetween string
	}

	toVisit := []pathState{
		{start, 0, ""},
	}
	visited := map[Coordinate]bool{
		start: true,
	}
	for len(toVisit) > 0 {
		visiting := toVisit[0]
		toVisit = toVisit[1:]

		val := tunnelMap[visiting.coordinate]
		keysToGetHere := visiting.keysInBetween

		if visiting.steps != 0 && unicode.IsLower(val) {
			//keyVal := keyAsIntExpr(val)
			//pattern := fmt.Sprintf("^[%s]{%d}$", keysToGetHere, len(keysToGetHere))
			adjacencyEntries[val] = adjacencyEntry{
				steps:         visiting.steps,
				keysInBetween: keysToGetHere,
			}

			if !strings.ContainsRune(keysToGetHere, val) {
				keysToGetHere += string(val)
			}
		}

		if unicode.IsUpper(val) {
			keyVal := unicode.ToLower(val)
			if !strings.ContainsRune(keysToGetHere, keyVal) {
				keysToGetHere += string(keyVal)
			}
		}

		for _, neighbor := range visiting.coordinate.Neighbors() {
			valAtNeighbor := tunnelMap[neighbor]
			if visited[neighbor] {
				continue
			}
			if valAtNeighbor == '#' {
				continue
			}

			visited[neighbor] = true
			toVisit = append(toVisit, pathState{neighbor, visiting.steps + 1, keysToGetHere})
		}
	}

	return adjacencyEntries
}

func keyAsIntExpr(key rune) int {
	return int(math.Pow(2, float64(key-'a')))
}

func hasRequiredKeys(keysPossessed, keysNeeded string) bool {
	for _, key := range keysNeeded {
		if !strings.ContainsRune(keysPossessed, key) {
			return false
		}
	}
	return true
}

func main() {
	tunnelMap := parseMap(adventutil.Parse(18))

	part1 := MinStepsToFindAllKeys(tunnelMap)
	fmt.Printf("Part 1: %d\n", part1)
}
