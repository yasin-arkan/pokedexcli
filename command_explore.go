package main

import "fmt"

func commandExplore(cfg *config, name string) error {

	fmt.Println("Exploring ", name, "...")

	locationsResp, err := cfg.pokeapiClient.ListLocationDetails(name)

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, encounter := range locationsResp.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}

	return nil

}
