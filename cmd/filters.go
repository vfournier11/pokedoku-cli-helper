package cmd

import (
	"fmt"
	"pokedoku/pkg/filter"
	"strings"

	"github.com/spf13/cobra"
)

// filtersCmd represents the filters command
// this command was simply made to debug some of the hard-coded filters
var filtersCmd = &cobra.Command{
	Use: "filters",
	Run: func(cmd *cobra.Command, args []string) {
		filterArg := args[0]
		fmt.Println("Filter: " + filterArg)
		filters := strings.Split(filterArg, "_AND_")

		allFilters, err := filter.NewPokemonFilters(filters)
		if err != nil {
			panic(err)
		}
		pokemonNames := allFilters.Apply()
		for _, pokemonName := range pokemonNames {
			fmt.Println(pokemonName)
		}
	},
}

func init() {
	rootCmd.AddCommand(filtersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filtersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filtersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
