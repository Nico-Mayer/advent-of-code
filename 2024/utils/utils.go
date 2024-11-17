package utils

import (
	"fmt"
	"log"
	"os"
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
