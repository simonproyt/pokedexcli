package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.nextLocationsURL != nil {
		url = *cfg.nextLocationsURL
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locationAreas LocationAreasResponse
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationAreas.Next
	cfg.previousLocationsURL = locationAreas.Previous

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *cfg.previousLocationsURL

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locationAreas LocationAreasResponse
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationAreas.Next
	cfg.previousLocationsURL = locationAreas.Previous

	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	return nil
}
