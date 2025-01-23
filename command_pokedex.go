package main

import "fmt"

func commandPokedex(cfg *config, name string) error {

	fmt.Println("Your pokedex:")

	for key, _ := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", key)
	}

	return nil

}
