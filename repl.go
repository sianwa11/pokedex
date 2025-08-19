package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := &config{}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
				err := command.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
}


func cleanInput(input string) []string {
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	words := strings.Fields(input)
	return words
}

type LocationResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type config struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous *string `json:"previous"`
	Results []LocationResult `json:"results"`
}

type cliCommand struct {
	name				string
	description		string
	callback			func(*config) error
}



func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: 			"exit",
			description: 	"Exit the Pokedex",
			callback: 		commandExit,
		},
		"help": {
			name: 			"help",
			description: 	"Display this help message",
			callback: 		commandHelp,
		},
		"map" : {
			name: 			"map",
			description: 	"Display the map of the Pokedex",
			callback: 		commandMap,
		},
		"mapb": {
			name: 			"mapb",
			description: 	"Display the previous map of the Pokedex",	
			callback: 		commandMapb,
		},
	}
}