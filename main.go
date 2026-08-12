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
	fmt.Println("Sup homie? Keeping up with your reading?")

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
	myCommands.register("reset", handlerReset)
	myCommands.register("users", handlerUsers)
	myCommands.register("agg", handlerAgg)
	myCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	myCommands.register("feeds", handlerListFeeds)
	myCommands.register("follow", middlewareLoggedIn(handlerFollow))
	myCommands.register("following", middlewareLoggedIn(handlerFollowing))

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
}
