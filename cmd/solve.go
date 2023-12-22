package cmd

import (
	"fmt"
	"github.com/mtslzr/pokeapi-go"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:  "solve",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("solve called")

		rowFilters := strings.Split(args[0], ",")
		fmt.Println("row:", rowFilters)
		colFilters := strings.Split(args[1], ",")
		fmt.Println("col:", colFilters)
		//var rowFilterObjects []PokemonFilter

		selectedPokemon := make(map[string]bool)
		var cellFilters = make([]string, len(rowFilters)*len(colFilters))
		for i, rowFilter := range rowFilters {
			currentRowFilter, err := PokemonFilterFactory(rowFilter)
			if err != nil {
				panic(err)
			}

			for j, colFilter := range colFilters {
				currentColFilter, err := PokemonFilterFactory(colFilter)
				if err != nil {
					panic(err)
				}
				pokemon := PokemonFilters{[]PokemonFilter{currentRowFilter, currentColFilter}}.Apply()
				for _, pokemonName := range pokemon {
					if !selectedPokemon[pokemonName] {
						selectedPokemon[pokemonName] = true
						cellFilters[i*len(colFilters)+j] = pokemonName
						break
					}
				}
			}
		}

		fmt.Println("Solved!")
		filterDescriptionPadding := 12
		maxLength := 0
		for _, cellFilter := range cellFilters {
			if len(cellFilter) > maxLength {
				maxLength = len(cellFilter)
			}
		}
		fmt.Printf("%-*s|", filterDescriptionPadding, "")
		for _, colFilter := range colFilters {
			fmt.Printf(" %-*s |", maxLength, colFilter)
		}
		fmt.Println()
		fmt.Println(strings.Repeat("-", filterDescriptionPadding+1+(maxLength+3)*len(colFilters)))
		for i := 0; i < len(colFilters); i++ {
			fmt.Printf(" %*s |", filterDescriptionPadding-2, rowFilters[i%len(rowFilters)])
			currentRow := cellFilters[i*len(colFilters) : (i+1)*len(colFilters)]
			for j := 0; j < len(currentRow); j++ {
				fmt.Printf(" %-*s |", maxLength, currentRow[j])
			}
			fmt.Println()
		}
	},
}

func PokemonFilterFactory(filterDescription string) (PokemonFilter, error) {
	filterSplit := strings.Split(filterDescription, ":")
	filterType := filterSplit[0]
	filterValue := filterSplit[1]

	switch filterType {
	case "type", "t":
		return PokemonTypeFilter{filterValue}, nil
	case "generation", "g":
		switch filterValue {
		case "kanto", "1":
			filterValue = "1"
		case "johto", "2":
			filterValue = "2"
		case "hoenn", "3":
			filterValue = "3"
		case "sinnoh", "4":
			filterValue = "4"
		case "unova", "5":
			filterValue = "5"
		case "kalos", "6":
			filterValue = "6"
		case "alola", "7":
			filterValue = "7"
		case "galar", "8":
			filterValue = "8"
		default:
			return nil, fmt.Errorf("invalid generation: %s", filterValue)
		}

		return PokemonGenerationFilter{generation: filterValue}, nil
	default:
		return nil, fmt.Errorf("invalid filter type: %s", filterType)
	}
}

type PokemonFilter interface {
	Apply() []string
}

type PokemonFilters struct {
	filters []PokemonFilter
}

func (f PokemonFilters) Apply() []string {
	var pokemonNamesMap = make(map[string]int)
	for _, filter := range f.filters {
		for _, pokemonName := range filter.Apply() {
			pokemonNamesMap[pokemonName]++
		}
	}

	var pokemonNames []string
	for pokemonName, count := range pokemonNamesMap {
		if count == len(f.filters) {
			pokemonNames = append(pokemonNames, pokemonName)
		}
	}
	sort.Strings(pokemonNames)
	return pokemonNames
}

type PokemonTypeFilter struct {
	typeID string
}

func (f PokemonTypeFilter) Apply() []string {
	typeResult, err := pokeapi.Type(f.typeID)
	if err != nil {
		fmt.Println(err)
	}

	var pokemonNames []string
	for _, pokemon := range typeResult.Pokemon {
		pokemonNames = append(pokemonNames, pokemon.Pokemon.Name)
	}
	sort.Strings(pokemonNames)
	return pokemonNames
}

type PokemonGenerationFilter struct {
	generation string
}

func (f PokemonGenerationFilter) Apply() []string {
	generationResult, err := pokeapi.Generation(f.generation)
	if err != nil {
		fmt.Println(err)
	}

	var pokemonNames []string
	for _, pokemon := range generationResult.PokemonSpecies {
		pokemonNames = append(pokemonNames, pokemon.Name)
	}
	sort.Strings(pokemonNames)
	return pokemonNames
}

func init() {
	rootCmd.AddCommand(solveCmd)
}
