package day02

import (
	"fmt"

	"github.com/nico-mayer/aoc-2024/utils"
)

func part01(data string) {
	reports := formatData(data)

	var saveReportsCount int
	for _, report := range reports {
		if report.isSafe(false) {
			saveReportsCount++
		}
	}

	fmt.Println(saveReportsCount)
}

func (r *Report) isSafe(withDampening bool) bool {
	levelsValid := func(levels []int) bool {
		if len(levels) < 2 {
			return true
		}

		prevOrder := utils.Sign(levels[1] - levels[0])

		for i := 1; i < len(levels); i++ {
			curLevel, prevLevel := levels[i], levels[i-1]
			order := utils.Sign(curLevel - prevLevel)

			if order != 0 && order != prevOrder {
				return false

			}
			distance := utils.PositivDistance(curLevel, prevLevel)

			if distance > 3 || distance < 1 {
				return false
			}
			prevOrder = order
		}
		return true
	}

	if !withDampening {
		return levelsValid(r.levels)
	}

	for i := 0; i < len(r.levels); i++ {
		lv := make([]int, len(r.levels)-1)
		copy(lv[:i], r.levels[:i])
		copy(lv[i:], r.levels[i+1:])
		if levelsValid(lv) {
			return true
		}
	}

	return false
}
