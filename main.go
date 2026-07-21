package main

import (
	"fmt"
	"log"

	"github.com/BredSandowich/rss_aggregator/internal/config"
)

func main() {
	fmt.Println("Sup homie?")

	configuration, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	currentState := state{Config : configuration}

	commandsMap := make(map[string]func(*state, command) error)
	myCommands := commands{registeredCommands : commandsMap}
	
	fmt.Println(configuration)

}