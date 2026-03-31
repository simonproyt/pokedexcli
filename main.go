package main

import (
	"bufio"
	"fmt"
	"os"
)

type config struct {
	nextLocationsURL     *string
	previousLocationsURL *string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}

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
