package day04

import (
	"fmt"
)

const (
	East = iota
	NorthEast
	North
	NorthWest
	West
	SouthWest
	South
	SouthEast
)

func part01(lines []string) {
	solution := 0
	for row, line := range lines {
		for col := range line {
			if checkNorth(row, col, lines, "XMAS") {
				solution++
				fmt.Printf("FOUND North at row: %d, col: %d\n", row, col)
			}
			if checkNorthEast(row, col, lines, "XMAS") {
				solution++
				fmt.Printf("FOUND NorthEast at row: %d, col: %d\n", row, col)
			}
			if checkEast(row, col, lines, "XMAS") {
				solution++
				fmt.Printf("FOUND East at row: %d, col: %d\n", row, col)
			}
			if checkSouthEast(row, col, lines, "XMAS") {
				solution++
				fmt.Printf("FOUND SouthEast at row: %d, col: %d\n", row, col)
			}
			if checkSouth(row, col, lines, "XMAS") {
				solution++
				fmt.Printf("FOUND South at row: %d, col: %d\n", row, col)
			}
			if checkSouthWest(row, col, lines, "XMAS") {
				solution++
				fmt.Printf("FOUND SouthWest at row: %d, col: %d\n", row, col)
			}
			if checkWest(row, col, lines, "XMAS") {
				solution++
				fmt.Printf("FOUND West at row: %d, col: %d\n", row, col)
			}
			if checkNorthWest(row, col, lines, "XMAS") {
				solution++
				fmt.Printf("FOUND NorthWest at row: %d, col: %d\n", row, col)
			}
		}
	}
	fmt.Printf("Solution: %d\n", solution)
}

func checkNorth(row int, col int, lines []string, targetWord string) bool {
	if targetWord[0] != lines[row][col] || !inBound(row, col, len(targetWord), lines, North) {
		return false
	}
	for i := 1; i < len(targetWord); i++ {
		if lines[row-i][col] != targetWord[i] {
			return false
		}
	}
	return true
}
func checkNorthEast(row int, col int, lines []string, targetWord string) bool {
	if targetWord[0] != lines[row][col] || !inBound(row, col, len(targetWord), lines, NorthEast) {
		return false
	}
	for i := 1; i < len(targetWord); i++ {
		if lines[row-i][col+i] != targetWord[i] {
			return false
		}
	}
	return true
}
func checkEast(row int, col int, lines []string, targetWord string) bool {
	if targetWord[0] != lines[row][col] || !inBound(row, col, len(targetWord), lines, East) {
		return false
	}
	for i := 1; i < len(targetWord); i++ {
		if lines[row][col+i] != targetWord[i] {
			return false
		}
	}
	return true
}
func checkSouthEast(row int, col int, lines []string, targetWord string) bool {
	if targetWord[0] != lines[row][col] || !inBound(row, col, len(targetWord), lines, SouthEast) {
		return false
	}
	for i := 1; i < len(targetWord); i++ {
		if lines[row+i][col+i] != targetWord[i] {
			return false
		}
	}
	return true
}
func checkSouth(row int, col int, lines []string, targetWord string) bool {
	if targetWord[0] != lines[row][col] || !inBound(row, col, len(targetWord), lines, South) {
		return false
	}
	for i := 1; i < len(targetWord); i++ {
		if lines[row+i][col] != targetWord[i] {
			return false
		}
	}
	return true
}
func checkSouthWest(row int, col int, lines []string, targetWord string) bool {
	if targetWord[0] != lines[row][col] || !inBound(row, col, len(targetWord), lines, SouthWest) {
		return false
	}
	for i := 1; i < len(targetWord); i++ {
		if lines[row+i][col-i] != targetWord[i] {
			return false
		}
	}
	return true
}
func checkWest(row int, col int, lines []string, targetWord string) bool {
	if targetWord[0] != lines[row][col] || !inBound(row, col, len(targetWord), lines, West) {
		return false
	}
	for i := 1; i < len(targetWord); i++ {
		if lines[row][col-i] != targetWord[i] {
			return false
		}
	}
	return true
}
func checkNorthWest(row int, col int, lines []string, targetWord string) bool {
	if targetWord[0] != lines[row][col] || !inBound(row, col, len(targetWord), lines, NorthWest) {
		return false
	}
	for i := 1; i < len(targetWord); i++ {
		if lines[row-i][col-i] != targetWord[i] {
			return false
		}
	}
	return true
}

func inBound(row, col, wordLength int, lines []string, direction int) bool {
	maxRows := len(lines)
	maxCols := len(lines[0])

	switch direction {
	case North:
		return row-wordLength+1 >= 0
	case NorthEast:
		return row-wordLength+1 >= 0 && col+wordLength <= maxCols
	case East:
		return col+wordLength <= maxCols
	case SouthEast:
		return row+wordLength <= maxRows && col+wordLength <= maxCols
	case South:
		return row+wordLength <= maxRows
	case SouthWest:
		return row+wordLength <= maxRows && col-wordLength+1 >= 0
	case West:
		return col-wordLength+1 >= 0
	case NorthWest:
		return row-wordLength+1 >= 0 && col-wordLength+1 >= 0
	default:
		return false
	}
}
