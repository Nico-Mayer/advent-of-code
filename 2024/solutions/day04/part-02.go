package day04

import (
	"fmt"

	"github.com/nico-mayer/aoc-2024/utils"
)

func part02(lines []string) {
	solution := 0
	for row, line := range lines {
		for col, r := range line {
			if r != rune('A') {
				continue
			}

			if crossInBound(row, col, lines, "MAS") {
				if checkCross(row, col, lines, "MAS") {
					solution++
				}
			}
		}
	}
	fmt.Printf("Solution: %d\n", solution)
}

func crossInBound(row int, col int, lines []string, targetWord string) bool {
	maxCols := len(lines[0]) - 1
	maxRows := len(lines) - 1
	checkLen := (len(targetWord) - 1) / 2

	if col+checkLen > maxCols || col-checkLen < 0 || row+checkLen > maxRows || row-checkLen < 0 {
		return false
	}

	return true
}

func checkCross(row int, col int, lines []string, targetWord string) bool {
	line1 := string([]rune{rune(lines[row+1][col-1]), rune(lines[row][col]), rune(lines[row-1][col+1])})
	line2 := string([]rune{rune(lines[row-1][col-1]), rune(lines[row][col]), rune(lines[row+1][col+1])})

	line1Valid := line1 == targetWord || line1 == utils.ReverseString(targetWord)
	line2Valid := line2 == targetWord || line2 == utils.ReverseString(targetWord)

	if line1Valid && line2Valid {
		return true
	}
	return false
}
