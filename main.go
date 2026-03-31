package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/simonproyt/pokedexcli/internal/pokecache"
)

type config struct {
	nextLocationsURL     *string
	previousLocationsURL *string
	cache                *pokecache.Cache
	caughtPokemon        map[string]Pokemon
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		cache:         pokecache.NewCache(5 * time.Minute),
		caughtPokemon: make(map[string]Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}
