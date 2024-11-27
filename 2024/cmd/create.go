/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [day]",
	Short: "Generates the daily boilerplate and retrieves Advent of Code (AOC) data for the specified day.",
	Long: `Generates boilerplate code for the specified day and fetches the corresponding user data from the Advent of Code website.
	To enable this functionality, ensure your sessionid is set in a .env file.
	Refer to the README for detailed setup instructions.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		day, err := strconv.Atoi(args[0])
		if err != nil || day < 1 || day > 24 {
			fmt.Println("Please provide a valid day between 1 and 24")
			return
		}
		fmt.Printf("Creating boilerplate for day %d\n", day)
		err = createDayBoilerplate(day)
		if err != nil {
			fmt.Printf("ERROR: creating boilerplate for day %d\n", day)
			return
		}
		fmt.Printf("Successfully created boilerplate for day %d\n", day)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func createDayBoilerplate(day int) error {
	// Define the folder and file paths
	dayFolder := fmt.Sprintf("./solutions/day%02d", day)
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
		inputContent := fetchDayData(day)

		if err := os.WriteFile(inputFile, []byte(inputContent), 0644); err != nil {
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
		codeContent := fmt.Sprintf(`package solutions

import (
	"github.com/nico-mayer/aoc-2024/utils"
	"fmt"
)

func Dai%02d() {
	data := utils.LoadData(%d, true)
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
		codeContent := fmt.Sprintf(`package solutions

import (
	"fmt"
)

func part01(data string) {
	fmt.Println(data)
}`)

		if err := os.WriteFile(part01File, []byte(codeContent), 0644); err != nil {
			return fmt.Errorf("failed to create code file: %v", err)
		}
	}

	// Create part-02 file
	if _, err := os.Stat(part02File); os.IsNotExist(err) {
		codeContent := fmt.Sprintf(`package solutions

import (
	"fmt"
)

func part02(data string) {
	fmt.Println(data)
}`)

		if err := os.WriteFile(part02File, []byte(codeContent), 0644); err != nil {
			return fmt.Errorf("failed to create code file: %v", err)
		}
	}

	return nil
}

func fetchDayData(day int) string {
	year := os.Getenv("AOC_YEAR")
	sessionCookie := os.Getenv("SESSION_COOKIE")
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%d/input", year, day)
	errorPrefix := "ERROR (Fetching AOC Data): "

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("%s creating request: %v\n", errorPrefix, err)
		return ""
	}

	// Add cookies using the http.Cookie struct
	cookies := []*http.Cookie{
		{Name: "session", Value: sessionCookie},
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s making request: %v\n", errorPrefix, err)
		return ""
	}
	defer resp.Body.Close() // Ensure the response body is closed

	// Check if the status code is OK (200)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s received status code: %d\n", errorPrefix, resp.StatusCode)
		return ""
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s reading response body: %v\n", errorPrefix, err)
		return ""
	}

	return string(body)
}
