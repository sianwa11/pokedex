package main

import (
	"time"

	"github.com/sianwa11/pokedex/internal/pokecache"
)


var (
	PokeCache *pokecache.Cache
	LocationURL string
	PokemonURL string
	MyPokedex *Pokedex
)

func main() {
	PokeCache = pokecache.NewCache(5 * time.Minute)
	LocationURL = "https://pokeapi.co/api/v2/location-area/"
	PokemonURL =  "https://pokeapi.co/api/v2/pokemon/"

	MyPokedex = &Pokedex{
		caught: make(map[string]Pokemon),
	}

	startRepl()
}