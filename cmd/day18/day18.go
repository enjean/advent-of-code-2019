package main

import (
	"container/heap"
	"fmt"
	. "github.com/enjean/advent-of-code-2019/internal/adventutil/coordinate"
	"regexp"
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
	key          rune
	keysFound    string
	steps        int
	index        int
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
	key           rune
	steps         int
	keysInBetween string
}

func MinStepsToFindAllKeys(tunnelMap map[Coordinate]rune) int {
	var currentCoord Coordinate
	numKeys := 0
	adjacencyMatrix := make(map[rune][]adjacencyEntry)
	for coord, val := range tunnelMap {
		if val == '@' {
			currentCoord = coord
			continue
		}
		if unicode.IsLower(val) {
			numKeys++
			adjacencyMatrix[val] = buildAdjacencyEntries(coord, tunnelMap)
		}
	}

	//for key, entries := range adjacencyMatrix {
	//	fmt.Println(string(key) + ":")
	//	for _, entry := range entries {
	//		fmt.Printf("  %s %d %s\n", string(entry.key), entry.steps, entry.keysInBetween)
	//	}
	//}

	var pq PriorityQueue

	for _, reachableKey := range reachableKeys(tunnelMap, currentCoord, nil) {
		heap.Push(&pq, &searchState{
			key:       reachableKey.key,
			keysFound: string(reachableKey.key),
			steps:     reachableKey.steps,
		})
	}

	for {
		processing := heap.Pop(&pq).(*searchState)
		//fmt.Printf("%d\n", pq.Len())
		//fmt.Printf("Processing %s %d\n", processing.keysFound, processing.steps)
		if len(processing.keysFound) == numKeys {
			return processing.steps
		}

		//reachableKeys := reachableKeys(tunnelMap, processing.coordinate, processing.keysFound)
		//fmt.Printf("Reachable keys = %v\n", reachableKeys)
		for _, adjacencyEntry := range adjacencyMatrix[processing.key] {
			if strings.ContainsRune(processing.keysFound, adjacencyEntry.key) {
				continue
			}
			matchExpr := fmt.Sprintf("^[%s]{%d}$", processing.keysFound, len(adjacencyEntry.keysInBetween))
			match, _ := regexp.MatchString(matchExpr, adjacencyEntry.keysInBetween)
			if match {
				//fmt.Printf("Can reach %s\n", string(adjacencyEntry.key))
				heap.Push(&pq, &searchState{
					key: adjacencyEntry.key,
					keysFound:  processing.keysFound + string(adjacencyEntry.key),
					steps:      processing.steps + adjacencyEntry.steps,
				})
			}
		}
	}
}

func buildAdjacencyEntries(start Coordinate, tunnelMap map[Coordinate]rune) []adjacencyEntry {
	var adjacencyEntries []adjacencyEntry

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

		if visiting.steps !=0 && unicode.IsLower(val) {
			//pattern := fmt.Sprintf("^[%s]{%d}$", keysToGetHere, len(keysToGetHere))
			adjacencyEntries = append(adjacencyEntries, adjacencyEntry{
				key:           val,
				steps:         visiting.steps,
				keysInBetween: keysToGetHere,
			})

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

func appendKeysFound(keysFound map[rune]bool, key rune) map[rune]bool {
	newKeysFound := make(map[rune]bool)
	for k, v := range keysFound {
		newKeysFound[k] = v
	}
	newKeysFound[key] = true
	return newKeysFound
}

func appendKeyFindOrder(keyFindOrder []rune, key rune) []rune {
	var newKeyFindOrder []rune
	newKeyFindOrder = append(newKeyFindOrder, keyFindOrder...)
	newKeyFindOrder = append(newKeyFindOrder, key)
	return newKeyFindOrder
}

type reachableKey struct {
	key   rune
	coord Coordinate
	steps int
}

func reachableKeys(tunnelMap map[Coordinate]rune, start Coordinate, keysFound map[rune]bool) []reachableKey {
	var reachableKeys []reachableKey

	type coordSteps struct {
		coordinate Coordinate
		steps      int
	}

	toVisit := []coordSteps{
		{start, 0},
	}
	visited := map[Coordinate]bool{
		start: true,
	}
	for len(toVisit) > 0 {
		visiting := toVisit[0]
		toVisit = toVisit[1:]

		for _, neighbor := range visiting.coordinate.Neighbors() {
			valAtNeighbor := tunnelMap[neighbor]
			if visited[neighbor] {
				continue
			}
			if valAtNeighbor == '#' {
				continue
			}
			if unicode.IsLower(valAtNeighbor) && !keysFound[valAtNeighbor] {
				reachableKeys = append(reachableKeys, reachableKey{
					key:   valAtNeighbor,
					coord: neighbor,
					steps: visiting.steps + 1,
				})
				continue
			}
			if unicode.IsUpper(valAtNeighbor) && !keysFound[unicode.ToLower(valAtNeighbor)] {
				continue
			}
			visited[neighbor] = true
			toVisit = append(toVisit, coordSteps{neighbor, visiting.steps + 1})
		}
	}

	return reachableKeys
}

func main() {

}
