package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nico-mayer/aoc-2024/utils"
)

type Rule struct {
	X int
	Y int
}

type Update struct {
	Pages []int
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

func formatInput(lines []string) (rules *utils.Set[Rule], updates []Update) {
	rules = utils.NewSet[Rule]()
	insideRulesBlock := true
	for _, line := range lines {
		if len(line) == 0 {
			insideRulesBlock = false
			continue
		}

		if insideRulesBlock {
			split := strings.Split(line, "|")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			rule := Rule{X: x, Y: y}
			rules.Add(rule)
		} else {
			split := strings.Split(line, ",")
			var update Update
			for _, pageNrStr := range split {
				pageNr, _ := strconv.Atoi(pageNrStr)
				update.Pages = append(update.Pages, pageNr)
			}
			updates = append(updates, update)
		}
	}
	return
}
