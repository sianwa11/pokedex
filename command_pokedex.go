package main

import "fmt"

func commandPokedex(_ *config, _ string) error {

	for _, p := range MyPokedex.caught {
		fmt.Printf(" -%v\n", p.Name)
	}

	return nil
}