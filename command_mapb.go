package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cfg *config, flag string) error {
	url := cfg.Previous

	if url == nil {
		return fmt.Errorf("no previous URL available")
	}

	if cached, ok := PokeCache.Get(*url); ok {
		if err := json.Unmarshal(cached, &cfg); err != nil {
			return fmt.Errorf("error unmarshalling location: %w", err)
		}

		for _, result := range cfg.Results {
			fmt.Println(result.Name)
		}
		return nil
	}

	res, err := http.Get(*url)
	if err != nil {
		return fmt.Errorf("error getting response: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	
	PokeCache.Add(*url, body)

	if err := json.Unmarshal(body, &cfg); err != nil {
		return fmt.Errorf("error unmarshalling location: %w", err)
	}


	for _, result := range cfg.Results {
		fmt.Println(result.Name)
	}

	return nil
}