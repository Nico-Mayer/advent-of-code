package cmd

import (
	"fmt"
	"strconv"

	"github.com/nico-mayer/aoc-2024/solutions"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [day]",
	Short: "Runs the solution for the specified day",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		day, err := strconv.Atoi(args[0])
		if err != nil || day < 1 || day > 24 {
			fmt.Println("Provide a valid number between 1 and 24")
			return
		}
		callback, exists := solutions.Solutions[day]
		useTestInputData, _ := cmd.Flags().GetBool("test")
		partToExecute, _ := cmd.Flags().GetInt("part")

		if partToExecute != 0 && partToExecute != 1 && partToExecute != 2 {
			fmt.Println("Invalid part specified, choose between 1 and 2 or leave empty to execute both")
			return
		}
		if exists {
			callback(useTestInputData, day, partToExecute)
		} else {
			fmt.Printf("There is no Run function for day %d, check if you added it to solutions map\n", day)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	runCmd.Flags().BoolP("test", "t", false, "Specify if this day should use the test dataset")
	runCmd.Flags().IntP("part", "p", 0, "Specify witch part to execute")
}
