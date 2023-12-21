package cmd

import (
	"fmt"
	"github.com/mtslzr/pokeapi-go"
	"sort"

	"github.com/spf13/cobra"
)

// findByAbilityCmd represents the findByAbility command
var findByAbilityCmd = &cobra.Command{
	Use:  "find-by-ability",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		abilityID := args[0]
		ability, err := pokeapi.Ability(abilityID)
		if err != nil {
			fmt.Println(err)
		}
		var pokemonNames []string
		for _, pokemon := range ability.Pokemon {
			pokemonNames = append(pokemonNames, pokemon.Pokemon.Name)
		}
		sort.Strings(pokemonNames)
		for _, pokemonName := range pokemonNames {
			fmt.Println(pokemonName)
		}
	},
}

func init() {
	rootCmd.AddCommand(findByAbilityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findByAbilityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findByAbilityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
