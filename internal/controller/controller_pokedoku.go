package controller

import (
	"pokedoku/internal/domain"
	"pokedoku/pkg/filter"
)

func ListPokemon(filters []string) ([]string, error) {
	allFilters, err := filter.NewPokemonFilters(filters)
	if err != nil {
		return nil, err
	}
	pokemonNames := allFilters.Apply()
	return pokemonNames, nil
}

func SolvePuzzle(columnFilters, rowFilters []string) ([][]string, error) {
	puzzleSolver := domain.PuzzleSolver{columnFilters, rowFilters}
	return puzzleSolver.Solve()
}
