package adventutil

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Parse(day int) []string {
	file, err := os.Open(fmt.Sprintf("cmd/day%d/input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Pow10(n int) int64 {
	result := int64(1)
	for i := 1; i <= n; i++ {
		result *= 10
	}
	return result
}

const MaxUint = ^uint(0)
const MinUint = 0

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1