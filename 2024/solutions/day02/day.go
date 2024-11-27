package solutions

import (
	"github.com/nico-mayer/aoc-2024/utils"
	"fmt"
)

func Dai02() {
	data := utils.LoadData(2, true)
	fmt.Println("Day 2:")
	fmt.Println("Part 01:")
	part01(data)
	fmt.Println()
	fmt.Println("Part 02:")
	part02(data)
}