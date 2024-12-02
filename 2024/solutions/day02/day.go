package day02

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc-2024/utils"
)

type Report struct {
	levels []int
}

func Run(test bool, day int, part int) {
	data := utils.LoadData(day, test)

	switch part {
	case 1:
		printHeadline(day, part)
		part01(data)
	case 2:
		printHeadline(day, part)
		part02(data)
	default:
		printHeadline(day, 1)
		part01(data)
		fmt.Println()
		printHeadline(day, 2)
		part02(data)
	}
}

func printHeadline(day, partNum int) {
	fmt.Printf("Day %d:\nPart %02d:\n", day, partNum)
}

func formatData(data string) []Report {
	var res []string
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if line == "" {

			continue
		}
		res = append(res, line)
	}

	var reports []Report
	for _, line := range lines {
		levels := strings.Fields(line)
		var report Report
		for _, level := range levels {
			v, err := strconv.Atoi(level)
			if err != nil {
				log.Fatalf("Error converting level string to int %v", err)
			}
			report.levels = append(report.levels, v)
		}
		reports = append(reports, report)
	}
	return reports
}
