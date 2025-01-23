package main

import (
	"fmt"
)

func commandInspect(cfg *config, name string) error {

	pokemon, ok := cfg.caughtPokemon[name]

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Println("Stats:")
	fmt.Printf("  -hp: %d\n", pokemon.Stats[0].BaseStat)
	fmt.Printf("  -attack: %d\n", pokemon.Stats[1].BaseStat)
	fmt.Printf("  -defense: %d\n", pokemon.Stats[2].BaseStat)
	fmt.Printf("  -special-attack: %d\n", pokemon.Stats[3].BaseStat)
	fmt.Printf("  -special-defense: %d\n", pokemon.Stats[4].BaseStat)
	fmt.Printf("  -speed: %d\n", pokemon.Stats[5].BaseStat)
	fmt.Println("Types:")
	for _, v := range pokemon.Types {
		fmt.Printf("  -%s\n", v.Type.Name)
	}

	return nil
}
