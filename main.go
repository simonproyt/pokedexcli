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
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		cache: pokecache.NewCache(5 * time.Minute),
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
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
