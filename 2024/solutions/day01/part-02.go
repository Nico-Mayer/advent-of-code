package day01

import "fmt"

func part02(data string) {
	left, right := formatData(data)
	var simScore int

	for _, leftV := range left {
		var occurence int
		for _, rightV := range right {
			if leftV == rightV {
				occurence++
			}
		}
		simScore += leftV * occurence
	}
	fmt.Println("Solution:", simScore)
}
