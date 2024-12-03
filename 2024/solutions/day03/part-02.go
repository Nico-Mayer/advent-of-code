package day03

import (
	"fmt"
	"regexp"
	"strconv"
)

func part02(lines []string) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	// Find all
	var matches [][]string
	for _, line := range lines {
		matches = append(matches, re.FindAllStringSubmatch(line, -1)...)
	}

	// Solution
	var solution int
	canMultiply := true
	for _, match := range matches {

		if len(match) > 2 && match[1] != "" && match[2] != "" && canMultiply {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			solution += x * y
		} else if match[0] == "do()" {
			canMultiply = true
		} else if match[0] == "don't()" {
			canMultiply = false
		}
	}
	fmt.Printf("Solution: %d\n", solution)
}
