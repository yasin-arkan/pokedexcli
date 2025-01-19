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

var commands map[string]cliCommand

func startRepl() {

	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(("Pokedex > "))
		scanner.Scan()
		input := scanner.Text()

		cleaned := cleanInput(input)

		for _, word := range cleaned {
			command, ok := commands[word]

			if ok {
				err := command.callback()
				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				}

				continue
			}

			fmt.Println("Unknown command")

		}

	}
}

func cleanInput(text string) []string {

	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")

	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}
