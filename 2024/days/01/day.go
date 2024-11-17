package day01

import (
	"aoc-2024/utils"
	"fmt"
)

func Run() {
	data := utils.LoadData(1, false)
	fmt.Println("Day 1:")
	fmt.Println("Part 01:")
	part01(data)
	fmt.Println()
	fmt.Println("Part 02:")
	part02(data)
}