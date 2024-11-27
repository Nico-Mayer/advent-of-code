package day02

import (
	"github.com/nico-mayer/aoc-2024/utils"
	"fmt"
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