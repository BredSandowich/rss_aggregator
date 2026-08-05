package main

import (
	"context"
	"fmt"
	"time"
	"errors"

	"github.com/google/uuid"
	"github.com/BredSandowich/rss_aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	//Check for missing arguments
	if len(cmd.Args) != 2 {
		return errors.New("Incorrect arguments entered")
	}
	
	//Get current user
	user, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url: 		cmd.Args[1]	,
		UserID: 	user.ID,
		})
	if err != nil {
		return err
}	

	fmt.Printf("Feed created: %v\n", feed.Name)
	fmt.Printf("From: %v\n", feed.Url)

	return nil
}