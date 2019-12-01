package adventutil

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Parse(day int) []string {
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", day))
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
