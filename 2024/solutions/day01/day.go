package day01

import (
	"fmt"
	"strings"

	"github.com/nico-mayer/aoc-2024/utils"
)

func Run(test bool, day int, part int) {
	lines := utils.LoadData(day, test)

	switch part {
	case 1:
		printHeadline(day, part)
		part01(lines)
	case 2:
		printHeadline(day, part)
		part02(lines)
	default:
		printHeadline(day, 1)
		part01(lines)
		fmt.Println()
		printHeadline(day, 2)
		part02(lines)
	}
}

func printHeadline(day, partNum int) {
	fmt.Printf("Day %d:\nPart %02d:\n", day, partNum)
}

func formatData(lines []string) (left, right []int) {
	for _, line := range lines {
		splitLine := strings.Fields(line)

		left = append(left, utils.ParseInt(splitLine[0]))
		right = append(right, utils.ParseInt(splitLine[1]))
	}
	return left, right
}
