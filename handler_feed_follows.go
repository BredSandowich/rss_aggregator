package main

import (
	"context"
	"fmt"
	"time"
	"errors"

	"github.com/google/uuid"
	"github.com/BredSandowich/rss_aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	//Check for missing arguments
	if len(cmd.Args) != 1 {
		return errors.New("Incorrect arguments entered")
	}
	
	//Get current user - if not utilizing the database.User in function header
	//user, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
	//if err != nil {
	//	return err
	//}

	//Get feed by URL
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
}	

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:      user.ID,
		FeedID: 	feed.ID,
		})
	if err != nil {
		return err
	}

	fmt.Printf("Following: %v\n", follow.FeedName)
	fmt.Printf("By: %v\n", follow.UserName)

	return nil
}


func handlerFollowing(s *state, cmd command, user database.User) error {
	//Get current user
	//user, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
	//if err != nil {
	//	return err
	//}

	//Get feed follows for current user
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
}	
	
	for _, feed := range feeds {
		fmt.Printf("Following feed: %v\n", feed.FeedName)
	}

	return nil
}