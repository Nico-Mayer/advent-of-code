package day03

import (
	"fmt"
	"regexp"
	"strconv"
)

func part01(lines []string) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all
	var matches [][]string
	for _, line := range lines {
		matches = append(matches, re.FindAllStringSubmatch(line, -1)...)
	}

	// Solution
	var solution int
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		solution += x * y
	}
	fmt.Printf("Solution: %d\n", solution)
}
