package main

import (
	"aoc-2024/utils"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if len(os.Args) == 3 {
		subcommand := os.Args[1]
		dayArg := os.Args[2]

		if subcommand != "create" {
			log.Fatalf("Unknown subcommand: %s", subcommand)
		}
		day, err := strconv.Atoi(dayArg)
		if err != nil || day <= 0 {
			log.Fatalf("Invalid day number: %s", dayArg)
		}
		err = utils.CreateDayBoilerplate(day)
		if err != nil {
			log.Fatalf("Error generating boilerplate: %v", err)
		}
		fmt.Printf("Successfully created boilerplate for Day %02d\n", day)
		os.Exit(0)
	}

	// day01.Run()
}
