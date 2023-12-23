package filter

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

type PokemonMythicalFilter struct {
}

func (f PokemonMythicalFilter) Apply() []string {
	return []string{
		"mew",
		"celebi",
		"jirachi",
		"deoxys",
		"manaphy",
		"darkrai",
		"shaymin",
		"arceus",
		"victini",
		"keldeo",
		"meloetta",
		"genesect",
		"diancie",
		"hoopa",
		"volcanion",
		"magearna",
		"marshadow",
		"zeraora",
		"meltan",
		"melmetal",
	}
}

func PokemonFilterFactory(filterDescription string) (PokemonFilter, error) {
	filterSplit := strings.Split(filterDescription, ":")
	filterType := filterSplit[0]
	filterValue := filterSplit[1]

	switch filterType {
	case "custom", "c":
		switch filterValue {
		case "legendary", "l":
			return PokemonLegendaryFilter{}, nil
		case "mythical", "m":
			return PokemonMythicalFilter{}, nil
		case "beast", "b":
			return PokemonAbilityFilter{"beast-boost"}, nil
		default:
			return nil, fmt.Errorf("invalid custom filter: %s", filterValue)
		}
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

func NewPokemonFilters(filterDescriptions []string) (PokemonFilters, error) {
	var pokemonFilters []PokemonFilter
	for _, filterDescription := range filterDescriptions {
		filter, err := PokemonFilterFactory(filterDescription)
		if err != nil {
			return PokemonFilters{}, err
		}
		pokemonFilters = append(pokemonFilters, filter)
	}
	return PokemonFilters{pokemonFilters}, nil
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

type PokemonLegendaryFilter struct {
}

func (f PokemonLegendaryFilter) Apply() []string {
	return []string{
		"articuno",
		"zapdos",
		"moltres",
		"mewtwo",
		"raikou",
		"entei",
		"suicune",
		"lugia",
		"ho-oh",
		"regirock",
		"regice",
		"registeel",
		"latias",
		"latios",
		"kyogre",
		"groudon",
		"rayquaza",
		"uxie",
		"mesprit",
		"azelf",
		"dialga",
		"palkia",
		"heatran",
		"regigigas",
		"giratina",
		"cresselia",
		"cobalion",
		"terrakion",
		"virizion",
		"tornadus",
		"thundurus",
		"reshiram",
		"zekrom",
		"landorus",
		"kyurem",
		"xerneas",
		"yveltal",
		"zygarde",
		"tapu-koko",
		"tapu-lele",
		"tapu-bulu",
		"tapu-fini",
		"cosmog",
		"cosmoem",
		"solgaleo",
		"lunala",
		"necrozma",
		"magearna",
		"marshadow",
	}
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

type PokemonAbilityFilter struct {
	abilityID string
}

func (f PokemonAbilityFilter) Apply() []string {
	abilityResult, err := pokeapi.Ability(f.abilityID)
	if err != nil {
		fmt.Println(err)
	}

	var pokemonNames []string
	for _, pokemon := range abilityResult.Pokemon {
		pokemonNames = append(pokemonNames, pokemon.Pokemon.Name)
	}
	sort.Strings(pokemonNames)
	return pokemonNames
}
