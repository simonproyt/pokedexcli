package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("catch requires a pokemon name")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	url := "https://pokeapi.co/api/v2/pokemon/" + name

	var body []byte

	if val, ok := cfg.cache.Get(url); ok {
		body = val
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return fmt.Errorf("failed to fetch pokemon: %d", res.StatusCode)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(url, body)
	}

	var pokemon Pokemon
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
