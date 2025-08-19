package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}


func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Print("\n")
	fmt.Print("\n")
	// fmt.Println("help: Displays a help message")
	// fmt.Println("exit: Exit the Pokedex")
	return nil
}

func main() {
	commands := map[string]cliCommand {
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
	}

	userInput := ""

	scanner := bufio.NewScanner(os.Stdin);
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		userInput = scanner.Text()

		userInput = strings.TrimSpace(strings.ToLower(userInput))

		switch userInput { 
			case "exit":
				commands["exit"].callback()
			case "help":
				commands["help"].callback()
				fmt.Printf("%s: %s\n",commands["help"].name, commands["help"].description)
				fmt.Printf("%s: %s\n",commands["exit"].name, commands["exit"].description)
			default:
				fmt.Printf("Unknown command: %s\n", userInput)
		}
		
	}
}