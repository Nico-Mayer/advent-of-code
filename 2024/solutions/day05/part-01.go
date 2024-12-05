package day05

import (
	"fmt"

	"github.com/nico-mayer/aoc-2024/utils"
)

func part01(lines []string) {
	rules, updates := formatInput(lines)
	solution := 0

	for _, update := range updates {
		valid := update.isValid(rules)

		if valid {
			halfIndex := len(update.Pages) / 2
			solution += update.Pages[halfIndex]
		}
	}

	fmt.Println(solution)
}

func (u *Update) isValid(rules *utils.Set[Rule]) bool {
	valid := true

	rules.ForEach(func(rule Rule) {
		indexOfX, xExists := utils.FindIndex(u.Pages, rule.X)
		if !xExists {
			return
		}
		indexOfY, yExists := utils.FindIndex(u.Pages, rule.Y)
		if !yExists {
			return
		}
		if indexOfX > indexOfY {
			valid = false
			return
		}
	})

	return valid
}

func (u *Update) sort(rules *utils.Set[Rule]) {
	for {
		isSorted := true
		for i := 0; i < len(u.Pages)-1; i++ {
			testRule := Rule{
				X: u.Pages[i],
				Y: u.Pages[i+1],
			}
			if !rules.Contains(testRule) {
				x := u.Pages[i]
				y := u.Pages[i+1]
				u.Pages[i] = y
				u.Pages[i+1] = x
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}
