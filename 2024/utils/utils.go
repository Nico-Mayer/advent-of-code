package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func LoadData(day int, test bool) string {
	filePath := fmt.Sprintf("./days/%02d/input.txt", day)
	if test {
		filePath = fmt.Sprintf("./days/%02d/input.test.txt", day)
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Panic("ERROR:", err)
	}

	return string(fileContent)
}

func CreateDayBoilerplate(day int) error {
	// Define the folder and file paths
	dayFolder := fmt.Sprintf("./days/%02d", day)
	inputFile := filepath.Join(dayFolder, "input.txt")
	testInputFile := filepath.Join(dayFolder, "input.test.txt")
	codeFile := filepath.Join(dayFolder, "day.go")
	part01File := filepath.Join(dayFolder, "part-01.go")
	part02File := filepath.Join(dayFolder, "part-02.go")

	// Create the day folder
	if err := os.MkdirAll(dayFolder, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create folder: %v", err)
	}

	// Create the input files
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		if err := os.WriteFile(inputFile, []byte(""), 0644); err != nil {
			return fmt.Errorf("failed to create input file: %v", err)
		}
	}
	if _, err := os.Stat(testInputFile); os.IsNotExist(err) {
		if err := os.WriteFile(testInputFile, []byte(""), 0644); err != nil {
			return fmt.Errorf("failed to create input file: %v", err)
		}
	}

	// Create the day.go file
	if _, err := os.Stat(codeFile); os.IsNotExist(err) {
		codeContent := fmt.Sprintf(`package day%02d

import (
	"aoc-2024/utils"
	"fmt"
)

func Run() {
	data := utils.LoadData(%d, false)
	fmt.Println("Day %d:")
	fmt.Println("Part 01:")
	part01(data)
	fmt.Println()
	fmt.Println("Part 02:")
	part02(data)
}`, day, day, day)

		if err := os.WriteFile(codeFile, []byte(codeContent), 0644); err != nil {
			return fmt.Errorf("failed to create code file: %v", err)
		}
	}

	// Create part-01 file
	if _, err := os.Stat(part01File); os.IsNotExist(err) {
		codeContent := fmt.Sprintf(`package day%02d

import (
	"fmt"
)

func part01(data string) {
	fmt.Println(data)
}`, day)

		if err := os.WriteFile(part01File, []byte(codeContent), 0644); err != nil {
			return fmt.Errorf("failed to create code file: %v", err)
		}
	}

	// Create part-02 file
	if _, err := os.Stat(part02File); os.IsNotExist(err) {
		codeContent := fmt.Sprintf(`package day%02d

import (
	"fmt"
)

func part02(data string) {
	fmt.Println(data)
}`, day)

		if err := os.WriteFile(part02File, []byte(codeContent), 0644); err != nil {
			return fmt.Errorf("failed to create code file: %v", err)
		}
	}

	return nil
}
