package day04

import (
	"fmt"
)

type Direction struct {
	x int
	y int
}

var (
	East      = Direction{x: 1, y: 0}
	NorthEast = Direction{x: 1, y: 1}
	North     = Direction{x: 0, y: 1}
	NorthWest = Direction{x: -1, y: 1}
	West      = Direction{x: -1, y: 0}
	SouthWest = Direction{x: -1, y: -1}
	South     = Direction{x: 0, y: -1}
	SouthEast = Direction{x: 1, y: -1}
)

func part01(lines []string) {
	solution := 0
	targetWord := "XMAS"
	directions := []Direction{
		North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest,
	}

	for row, line := range lines {
		for col := range line {
			for _, dir := range directions {
				if checkDirection(row, col, lines, targetWord, dir) {
					solution++
				}
			}
		}
	}

	fmt.Printf("Solution: %d\n", solution)
}

func checkDirection(row int, col int, lines []string, targetWord string, direction Direction) bool {
	if targetWord[0] != lines[row][col] || !inBound(row, col, len(targetWord), lines, direction) {
		return false
	}
	for i := 1; i < len(targetWord); i++ {
		if lines[row+(i*direction.y)][col+(i*direction.x)] != targetWord[i] {
			return false
		}
	}
	return true
}

func inBound(row, col, wordLength int, lines []string, direction Direction) bool {
	maxRows := len(lines)
	maxCols := len(lines[0])

	xSafe := col+direction.x*(wordLength-1) >= 0 && col+direction.x*(wordLength-1) < maxCols
	ySafe := row+direction.y*(wordLength-1) >= 0 && row+direction.y*(wordLength-1) < maxRows

	return xSafe && ySafe
}
