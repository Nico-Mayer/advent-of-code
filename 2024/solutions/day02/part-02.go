package day02

import (
	"fmt"
)

func part02(lines []string) {
	reports := formatData(lines)

	var saveReportsCount int
	for _, report := range reports {
		if report.isSafe(true) {
			saveReportsCount++
		}
	}

	fmt.Println(saveReportsCount)
}
