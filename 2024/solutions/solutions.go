package solutions

import (
	"github.com/nico-mayer/aoc-2024/solutions/day01"
	"github.com/nico-mayer/aoc-2024/solutions/day02"
	"github.com/nico-mayer/aoc-2024/solutions/day03"
	"github.com/nico-mayer/aoc-2024/solutions/day04"
	"github.com/nico-mayer/aoc-2024/solutions/day05"
	"github.com/nico-mayer/aoc-2024/solutions/day06"
)

var Solutions = map[int]func(bool, int, int){
	1: day01.Run,
	2: day02.Run,
	3: day03.Run,
	4: day04.Run,
	5: day05.Run,
	6: day06.Run,
	// Add more days here
}
