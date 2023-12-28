package pokeapi

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const pokeApiGraphqlEndpoint = "https://beta.pokeapi.co/graphql/v1beta"

const pokeApiGraphqlQueryMegaForm = `{
  pokemon_v2_pokemonform(where: {is_mega: {_eq: true}}) {
    pokemon_v2_pokemon {
      name
    }
  }
}`

type PokemonResponse struct {
	Data struct {
		Pokemon_v2_pokemonform []struct {
			Pokemon_v2_pokemon struct {
				Name string `json:"name"`
			} `json:"pokemon_v2_pokemon"`
		} `json:"pokemon_v2_pokemonform"`
	} `json:"data"`
}

func Mega() ([]string, error) {
	httpClient := &http.Client{}

	// Create a map to hold the request body data
	requestBodyMap := map[string]string{
		"query": pokeApiGraphqlQueryMegaForm,
	}

	// Convert the map to JSON
	requestBodyJson, err := json.Marshal(requestBodyMap)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", pokeApiGraphqlEndpoint, bytes.NewBuffer(requestBodyJson))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Method-Used", "graphiql")

	res, err := httpClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Decode the JSON from the original response body
	res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	var pokemonResponse PokemonResponse
	err = json.NewDecoder(res.Body).Decode(&pokemonResponse)
	if err != nil {
		return nil, err
	}

	var pokemonNames []string
	for _, pokemon := range pokemonResponse.Data.Pokemon_v2_pokemonform {
		pokemonNames = append(pokemonNames, pokemon.Pokemon_v2_pokemon.Name)
	}

	return pokemonNames, nil
}
