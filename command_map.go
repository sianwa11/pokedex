package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config) error {
	
	url := cfg.Next

	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
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

	if err := json.Unmarshal(body, &cfg); err != nil {
		return err
	}

	for _, result := range cfg.Results {
		fmt.Println(result.Name)
	}
	return nil
}