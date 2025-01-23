package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, name string) error {

	pokemonResp, err := cfg.pokeapiClient.ListPokemon(name)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	baseExp := pokemonResp.BaseExperience

	if rand.Intn(baseExp) > 40 {
		fmt.Printf("%s was caught!\n", name)
		cfg.caughtPokemon[name] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil

}
