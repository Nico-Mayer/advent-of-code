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
	lines := utils.LoadData(day, test)

	switch part {
	case 1:
		printHeadline(day, part)
		part01(lines)
	case 2:
		printHeadline(day, part)
		part02(lines)
	default:
		printHeadline(day, 1)
		part01(lines)
		fmt.Println()
		printHeadline(day, 2)
		part02(lines)
	}
}

func printHeadline(day, partNum int) {
	fmt.Printf("Day %d:\nPart %02d:\n", day, partNum)
}

func formatData(lines []string) []Report {
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
