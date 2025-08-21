package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	BaseExperience int `json:"base_experience"`
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
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if rand.IntN(5) < 4 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	MyPokedex.AddPokemon(pokemon)
	fmt.Printf("%s was caught!\n", pokemon.Name)


	return nil
}