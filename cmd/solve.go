package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pokedoku/pkg/filter"
	"strings"
)

type PuzzleSolver struct {
	columnFilters []string
	rowFilters    []string
}

func (p PuzzleSolver) Solve() ([][]string, error) {
	selectedPokemon := make(map[string]bool)
	cellFilterNames := make([][]string, len(p.rowFilters))
	for i, rowFilter := range p.rowFilters {
		cellFilterNames[i] = make([]string, len(p.columnFilters))
		for j, colFilter := range p.columnFilters {
			cellFilterObjects, err := filter.NewPokemonFilters([]string{rowFilter, colFilter})
			if err != nil {
				return nil, err
			}
			pokemon := cellFilterObjects.Apply()
			for _, pokemonName := range pokemon {
				if !selectedPokemon[pokemonName] {
					selectedPokemon[pokemonName] = true
					cellFilterNames[i][j] = pokemonName
					break
				}
			}
		}
	}
	return cellFilterNames, nil
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

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:  "solve",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("solve called")

		columnFilters := strings.Split(args[0], ",")
		fmt.Println("row:", columnFilters)
		rowFilters := strings.Split(args[1], ",")
		fmt.Println("col:", rowFilters)

		puzzleSolver := PuzzleSolver{columnFilters, rowFilters}
		solution, err := puzzleSolver.Solve()
		if err != nil {
			panic(err)
		}

		fmt.Println("Solved!")
		PrintSolvedPuzzle(columnFilters, rowFilters, solution)
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)
}
