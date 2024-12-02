package day02

import (
	"fmt"
)

func part02(data string) {
	reports := formatData(data)

	var saveReportsCount int
	for _, report := range reports {
		if report.isSafe(true) {
			saveReportsCount++
		}
	}

	fmt.Println(saveReportsCount)
}
