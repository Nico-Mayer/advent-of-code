package day01

import (
	"fmt"
	"strings"

	"github.com/nico-mayer/aoc-2024/utils"
)

func Run(test bool, day int, part int) {
	data := utils.LoadData(day, test)

	switch part {
	case 1:
		printHeadline(day, part)
		part01(data)
	case 2:
		printHeadline(day, part)
		part02(data)
	default:
		printHeadline(day, 1)
		part01(data)
		fmt.Println()
		printHeadline(day, 2)
		part02(data)
	}
}

func printHeadline(day, partNum int) {
	fmt.Printf("Day %d:\nPart %02d:\n", day, partNum)
}

func formatData(data string) (left, right []int) {
	splitData := strings.Split(data, "\n")

	for _, line := range splitData {
		if line == "" {
			continue
		}
		splitLine := strings.Fields(line)

		left = append(left, utils.ParseInt(splitLine[0]))
		right = append(right, utils.ParseInt(splitLine[1]))
	}
	return left, right
}
