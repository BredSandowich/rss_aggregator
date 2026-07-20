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