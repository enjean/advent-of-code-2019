package main

import (
	"container/heap"
	"fmt"
	. "github.com/enjean/advent-of-code-2019/internal/adventutil/coordinate"
	"math"
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
	key           int
	steps         int
	keysInBetween int
}

func MinStepsToFindAllKeys(tunnelMap map[Coordinate]rune) int {

	var pq PriorityQueue
	foundAllVal := 0
	adjacencyMatrix := make(map[int][]adjacencyEntry)
	for coord, val := range tunnelMap {
		if val == '@' {
			adjacencyMatrix[0] = buildAdjacencyEntries(coord, 0, tunnelMap)
			heap.Push(&pq, &searchState{
				key:       0,
				keysFound: 0,
				steps:     0,
			})
		}
		if unicode.IsLower(val) {
			key := keyAsIntExpr(val)
			foundAllVal += key
			adjacencyMatrix[key] = buildAdjacencyEntries(coord, key, tunnelMap)
		}
	}

	//for key, entries := range adjacencyMatrix {
	//	fmt.Println(string(key) + ":")
	//	for _, entry := range entries {
	//		fmt.Printf("  %s %d %d\n", string(entry.key), entry.steps, entry.keysInBetween)
	//	}
	//}

	for {
		processing := heap.Pop(&pq).(*searchState)
		//fmt.Printf("%d\n", pq.Len())
		fmt.Printf("Processing %d\n", processing.steps)
		if processing.keysFound == foundAllVal {
			return processing.steps
		}

		//reachableKeys := reachableKeys(tunnelMap, processing.coordinate, processing.keysFound)
		//fmt.Printf("Reachable keys = %v\n", reachableKeys)
		for _, adjacencyEntry := range adjacencyMatrix[processing.key] {
			if hasRequiredKeys(processing.keysFound, adjacencyEntry.key) {
				// already have this key
				continue
			}
			if hasRequiredKeys(processing.keysFound, adjacencyEntry.keysInBetween) {
				heap.Push(&pq, &searchState{
					key:       adjacencyEntry.key,
					keysFound: processing.keysFound + adjacencyEntry.key,
					steps:     processing.steps + adjacencyEntry.steps,
				})
			}
		}
	}
}

func buildAdjacencyEntries(start Coordinate, startKeyExpr int, tunnelMap map[Coordinate]rune) []adjacencyEntry {
	var adjacencyEntries []adjacencyEntry

	type pathState struct {
		coordinate    Coordinate
		steps         int
		keysInBetween int
	}

	toVisit := []pathState{
		{start, 0, 0},
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
			keyVal := keyAsIntExpr(val)
			//pattern := fmt.Sprintf("^[%s]{%d}$", keysToGetHere, len(keysToGetHere))
			adjacencyEntries = append(adjacencyEntries, adjacencyEntry{
				key:           keyVal,
				steps:         visiting.steps,
				keysInBetween: keysToGetHere,
			})

			if !hasRequiredKeys(keysToGetHere, keyVal) {
				keysToGetHere += keyVal
			}
		}

		if unicode.IsUpper(val) {
			keyVal := keyAsIntExpr(unicode.ToLower(val))
			if !hasRequiredKeys(keysToGetHere, keyVal) {
				keysToGetHere += keyVal
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

func hasRequiredKeys(keysPossessed, keysNeeded int) bool {
	return keysPossessed&keysNeeded == keysNeeded
}

func main() {

}
