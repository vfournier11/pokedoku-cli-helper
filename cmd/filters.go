package cmd

import (
	"fmt"
	"log"

	"pokedoku/internal/controller"

	"github.com/spf13/cobra"
)

// filterCmd represents the filters command
// this command was simply made to debug some of the hard-coded filters
var filterCmd = &cobra.Command{
	Use:  "filter",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		names, err := controller.ListPokemon(args)
		if err != nil {
			log.Fatal(err)
		}

		for _, name := range names {
			fmt.Println(name)
		}
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)
}
