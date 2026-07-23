package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/BredSandowich/rss_aggregator/internal/config"
	"github.com/BredSandowich/rss_aggregator/internal/database"
)

type state struct {
	Config *config.Config
	db     *database.Queries
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	//Check for missing arguments
	if len(cmd.Args) == 0 {
		return errors.New("No arguments entered")
	}

	//Check if user already exists
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	//Call method on config pointer
	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	//Print confirmation of user input
	fmt.Printf("User has been set to %s\n", cmd.Args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	//Check for missing arguments
	if len(cmd.Args) == 0 {
		return errors.New("No arguments entered")
	}
	//Create user through s.db
	user, err := s.db.CreateUsers(context.Background(), database.CreateUsersParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return err
	}

	//Call method on config
	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	//Print confirmation
	log.Printf("Created user: %v", user)
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	currentState, exists := c.registeredCommands[cmd.Name]
	if !exists {
		return errors.New("Command does not exist")
	} else {
		result := currentState(s, cmd)
		return result
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
