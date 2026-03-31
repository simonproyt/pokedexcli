package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("explore requires a location area name")
	}

	name := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + name

	fmt.Printf("Exploring %s...\n", name)

	if val, ok := cfg.cache.Get(url); ok {
		return processLocationAreaResponse(val)
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("failed to fetch location area: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	cfg.cache.Add(url, body)

	return processLocationAreaResponse(body)
}

func processLocationAreaResponse(body []byte) error {
	var locationArea LocationAreaResponse
	err := json.Unmarshal(body, &locationArea)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
