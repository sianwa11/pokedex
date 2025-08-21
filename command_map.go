package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	pokecache "github.com/sianwa11/pokedex/internal"
)

var pokeCache *pokecache.Cache

func init() {
	pokeCache = pokecache.NewCache(5 * time.Minute)
}

func commandMap(cfg *config) error {
	
	url := cfg.Next

	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	if cached, ok := pokeCache.Get(url); ok {
		if err := json.Unmarshal(cached, &cfg); err != nil {
			return err
		}

		for _, result := range cfg.Results {
			fmt.Println(result.Name)
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

	pokeCache.Add(url, body)

	if err := json.Unmarshal(body, &cfg); err != nil {
		return err
	}

	for _, result := range cfg.Results {
		fmt.Println(result.Name)
	}
	return nil
}