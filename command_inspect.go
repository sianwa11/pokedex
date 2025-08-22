package main

import "fmt"

func commandInspect(_ *config, pokemon string) error {

	if pokemon == "" {
		return fmt.Errorf("please indicate the pokemon you want to inspect")
	}

	p, ok := MyPokedex.caught[pokemon]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %+v\n", p.Name)
	fmt.Printf("Height: %+v\n", p.Height)
	fmt.Printf("Weight: %+v\n", p.Weight)
	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Printf(" -%v : %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf(" -%v \n", t.Type.Name)
	}


	return nil
}