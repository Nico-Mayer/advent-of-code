package day01

import (
	"fmt"
	"sort"

	"github.com/nico-mayer/aoc-2024/utils"
)

func part01(lines []string) {
	left, right := formatData(lines)
	sort.Ints(left)
	sort.Ints(right)
	var distance int

	for i, value := range left {
		distance += utils.PositivDistance(value, right[i])
	}
	fmt.Println("Solution:", distance)
}
