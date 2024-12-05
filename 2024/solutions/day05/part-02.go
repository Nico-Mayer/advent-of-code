package day05

import (
	"fmt"
)

func part02(lines []string) {
	rules, updates := formatInput(lines)
	solution := 0

	for _, update := range updates {
		valid := update.isValid(rules)

		if valid {
			continue
		} else {
			update.sort(rules)
			halfIndex := len(update.Pages) / 2
			solution += update.Pages[halfIndex]
		}
	}

	fmt.Println(solution)
}
