package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonEncounter struct {
	Pokemon struct {
		Name string  `json:"name"`
		URL  string  `json:"url"`
	}
}

type LocationAreaResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func commandExplore(_ * config, areaName string) error {
	if areaName == "" {
		return fmt.Errorf("area name cannot be empty")
		
	}

	url := LocationURL + areaName


	var location LocationAreaResponse

	if cached, ok := PokeCache.Get(url); ok {
		if err := json.Unmarshal(cached, &location); err != nil {
			return err
		}

		for _, result := range location.PokemonEncounters {
			fmt.Printf("- %s\n", result.Pokemon.Name)
		}
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	PokeCache.Add(url, body)

	if err := json.Unmarshal(body, &location); err != nil {
		return err
	}

	for _, result := range location.PokemonEncounters {
		fmt.Printf("- %s\n", result.Pokemon.Name)
	}

	return nil
}