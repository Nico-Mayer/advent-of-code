package day01

import (
	"github.com/nico-mayer/aoc-2024/utils"
	"fmt"
)

func Run() {
	data := utils.LoadData(1, true)
	fmt.Println("Day 1:")
	fmt.Println("Part 01:")
	part01(data)
	fmt.Println()
	fmt.Println("Part 02:")
	part02(data)
}