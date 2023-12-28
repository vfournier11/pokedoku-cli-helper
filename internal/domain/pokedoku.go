package domain

import (
	"pokedoku/pkg/filter"
)

type PuzzleSolver struct {
	ColumnFilters []string
	RowFilters    []string
}

func (p PuzzleSolver) Solve() ([][]string, error) {
	selectedPokemon := make(map[string]bool)
	cellFilterNames := make([][]string, len(p.RowFilters))
	for i, rowFilter := range p.RowFilters {
		cellFilterNames[i] = make([]string, len(p.ColumnFilters))
		for j, colFilter := range p.ColumnFilters {
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
