package main

import (
	"time"

	"github.com/sianwa11/pokedex/internal/pokecache"
)


var PokeCache *pokecache.Cache

func main() {
	PokeCache = pokecache.NewCache(5 * time.Minute)
	startRepl()
}