package cmd

import (
	"fmt"
	"github.com/mtslzr/pokeapi-go"
	"sort"

	"github.com/spf13/cobra"
)

// findByTypeCmd represents the findByTypes command
var findByTypeCmd = &cobra.Command{
	Use:  "find-by-type",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		typeID := args[0]

		typeResult, err := pokeapi.Type(typeID)
		if err != nil {
			fmt.Println(err)
		}

		var pokemonNames []string
		for _, pokemon := range typeResult.Pokemon {
			pokemonNames = append(pokemonNames, pokemon.Pokemon.Name)
		}
		sort.Strings(pokemonNames)
		for _, pokemonName := range pokemonNames {
			fmt.Println(pokemonName)
		}
	},
}

func init() {
	rootCmd.AddCommand(findByTypeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findByTypeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findByTypeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
