package main

import (
	"fmt"
	"log"
	"errors"

	"github.com/BredSandowich/rss_aggregator/internal/config"
)

type state struct {
	Config *config.Config
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
		return errors.New("no arguments entered")
	}

	//Call method on config pointer
	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	
	//Print confirmation of user input
	fmt.Printf("User has been set to %s\n", cmd.Args[0])
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