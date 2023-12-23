package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pokedoku/pkg/filter"
	"strings"
)

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

		selectedPokemon := make(map[string]bool)
		var cellFilterNames = make([]string, len(rowFilters)*len(columnFilters))
		for i, rowFilter := range rowFilters {
			for j, colFilter := range columnFilters {
				cellFilterObjects, err := filter.NewPokemonFilters([]string{rowFilter, colFilter})
				if err != nil {
					panic(err)
				}
				pokemon := cellFilterObjects.Apply()
				for _, pokemonName := range pokemon {
					if !selectedPokemon[pokemonName] {
						selectedPokemon[pokemonName] = true
						cellFilterNames[i*len(columnFilters)+j] = pokemonName
						break
					}
				}
			}
		}

		fmt.Println("Solved!")
		filterDescriptionPadding := 16
		maxLength := filterDescriptionPadding
		for _, cellFilter := range cellFilterNames {
			if len(cellFilter) > maxLength {
				maxLength = len(cellFilter)
			}
		}
		fmt.Printf("%-*s|", filterDescriptionPadding, "")
		for _, colFilter := range columnFilters {
			fmt.Printf(" %-*s |", maxLength, colFilter)
		}
		fmt.Println()
		fmt.Println(strings.Repeat("-", filterDescriptionPadding+1+(maxLength+3)*len(columnFilters)))
		for i := 0; i < len(columnFilters); i++ {
			fmt.Printf(" %*s |", filterDescriptionPadding-2, rowFilters[i%len(rowFilters)])
			currentRow := cellFilterNames[i*len(columnFilters) : (i+1)*len(columnFilters)]
			for j := 0; j < len(currentRow); j++ {
				fmt.Printf(" %-*s |", maxLength, currentRow[j])
			}
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)
}
