package solutions

import (
	"github.com/nico-mayer/aoc-2024/solutions/day01"
	"github.com/nico-mayer/aoc-2024/solutions/day02"
)

var Solutions = map[int]func(bool, int, int){
	1: day01.Run,
	2: day02.Run,
	// Add more days here
}