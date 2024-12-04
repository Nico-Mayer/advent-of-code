package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadData(day int, test bool) []string {
	filePath := fmt.Sprintf("./solutions/day%02d/input.txt", day)
	if test {
		filePath = fmt.Sprintf("./solutions/day%02d/input.test.txt", day)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Panic("ERROR:", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		lines = append(lines, scanner.Text())
	}
	return lines
}

func LoadDataAsString(day int, test bool) string {
	filePath := fmt.Sprintf("./solutions/day%02d/input.txt", day)
	if test {
		filePath = fmt.Sprintf("./solutions/day%02d/input.test.txt", day)
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Panic("ERROR:", err)
	}

	return strings.Trim(string(fileContent), "\n")
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

func Sign(v int) int {
	if v < 0 {
		return -1
	} else if v > 0 {
		return 1
	}
	return 0
}
func ReverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}
