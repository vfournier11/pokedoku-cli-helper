package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"pokedoku/internal/controller"
	"strings"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:  "solve",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		columnFilters := strings.Split(args[0], ",")
		rowFilters := strings.Split(args[1], ",")

		solvePuzzle, err := controller.SolvePuzzle(columnFilters, rowFilters)
		if err != nil {
			log.Fatal(err)
		}

		PrintSolvedPuzzle(columnFilters, rowFilters, solvePuzzle)
	},
}

func PrintSolvedPuzzle(columnFilters, rowFilters []string, cellFilterNames [][]string) {
	filterDescriptionPadding := 16
	maxLength := filterDescriptionPadding
	for _, row := range cellFilterNames {
		for _, cellFilter := range row {
			if len(cellFilter) > maxLength {
				maxLength = len(cellFilter)
			}
		}
	}
	fmt.Printf("%-*s|", filterDescriptionPadding, "")
	for _, colFilter := range columnFilters {
		fmt.Printf(" %-*s |", maxLength, colFilter)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("-", filterDescriptionPadding+1+(maxLength+3)*len(columnFilters)))
	for i, row := range cellFilterNames {
		fmt.Printf(" %*s |", filterDescriptionPadding-2, rowFilters[i%len(rowFilters)])
		for _, cellFilter := range row {
			fmt.Printf(" %-*s |", maxLength, cellFilter)
		}
		fmt.Println()
	}
}

func init() {
	rootCmd.AddCommand(solveCmd)
}
