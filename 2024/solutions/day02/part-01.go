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
	if len(r.levels) < 2 {
		return true
	}

	var subReport Report

	prevOrder := utils.Sign(r.levels[1] - r.levels[0])

	for i := 1; i < len(r.levels); i++ {
		curLevel, prevLevel := r.levels[i], r.levels[i-1]
		order := utils.Sign(curLevel - prevLevel)

		if order != 0 && order != prevOrder {
			if withDampening {
				subReport.levels = append(r.levels[:i], r.levels[i+1:]...)
				break
			} else {
				return false
			}
		}
		distance := utils.PositivDistance(curLevel, prevLevel)

		if distance > 3 || distance < 1 {
			if withDampening {
				subReport.levels = append(r.levels[:i], r.levels[i+1:]...)
				break
			} else {
				return false
			}
		}
		prevOrder = order
	}

	if withDampening {
		return subReport.isSafe(false)
	}

	return true
}
