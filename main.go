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
	
	err = configuration.SetUser("Brent")
	if err != nil {
		log.Fatal(err)
	}

	newConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(newConfig)

}