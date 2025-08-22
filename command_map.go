package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



func commandMap(cfg *config, flag string) error {
	
	url := cfg.Next

	if url == "" {
		url = LocationURL
	}

	if cached, ok := PokeCache.Get(url); ok {
		if err := json.Unmarshal(cached, &cfg); err != nil {
			return fmt.Errorf("error unmarshalling location: %w", err)
		}

		for _, result := range cfg.Results {
			fmt.Println(result.Name)
		}
		return nil
	}
	
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error getting response: %w", err)
	}
	defer res.Body.Close()	

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading stream: %w", err)
	}

	PokeCache.Add(url, body)

	if err := json.Unmarshal(body, &cfg); err != nil {
		return fmt.Errorf("error unmarshalling location: %w", err)
	}

	for _, result := range cfg.Results {
		fmt.Println(result.Name)
	}
	return nil
}