package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
)

type PokemonStats struct {
	BaseStat int `json:"base_stat"`
	Effort 	 int `json:"effort"`
	Stat struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}

type PokemonTypes struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int `json:"base_experience"`
	Height         int `json:"height"`
	Weight         int `json:"weight"`
	Stats          []PokemonStats `json:"stats"`
	Types          []PokemonTypes `json:"types"`
}

type Pokedex struct {
	caught map[string]Pokemon
}

func (p *Pokedex) AddPokemon(pokemon Pokemon) {
	p.caught[pokemon.Name] = pokemon
}

func commandCatch(_ *config, pokemonName string) error {

	if pokemonName == "" {
		return fmt.Errorf("pokemon name cannot be empty")
	}

	url := PokemonURL + pokemonName

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error getting request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading stream: %w", err)
	}

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return fmt.Errorf("error unmarshalling pokemon: %w", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if rand.IntN(5) < 4 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	MyPokedex.AddPokemon(pokemon)
	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command")

	return nil
}