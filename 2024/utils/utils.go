package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func LoadData(day int, test bool) string {
	filePath := fmt.Sprintf("./solutions/day%02d/input.txt", day)
	if test {
		filePath = fmt.Sprintf("./solutions/day%02d/input.test.txt", day)
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Panic("ERROR:", err)
	}

	return string(fileContent)
}

func ToIntSlice(stringSlice []string) ([]int, error) {
	intSlice := make([]int, len(stringSlice))
	for i, s := range stringSlice {
		v, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("ERROR: converting string slice to int slice")
			return nil, err
		}
		intSlice[i] = v
	}
	return intSlice, nil
}

func ParseInt(value string) int {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Error converting '%s' to int: %v", value, err)
	}
	return parsed
}

func PositivDistance(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
