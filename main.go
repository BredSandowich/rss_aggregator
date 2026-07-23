package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/BredSandowich/rss_aggregator/internal/config"
	"github.com/BredSandowich/rss_aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Sup homie?")

	configuration, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	//Load database URL to config struct
	db, err := sql.Open("postgres", configuration.DbUrl)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)
	currentState := state{Config: &configuration, db: dbQueries}

	commandsMap := make(map[string]func(*state, command) error)
	myCommands := commands{registeredCommands: commandsMap}

	myCommands.register("login", handlerLogin)
	myCommands.register("register", handlerRegister)

	if len(os.Args) < 2 {
		fmt.Println("Command does not exist")
		os.Exit(1)
	} else {
		newCommand := command{Name: os.Args[1], Args: os.Args[2:]}
		err = myCommands.run(&currentState, newCommand)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
	}
	fmt.Println(configuration)
}
