/*
Copyright © 2024 Nico Mayer
*/
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nico-mayer/aoc-2024/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cmd.Execute()
}
