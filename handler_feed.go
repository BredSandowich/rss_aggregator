package main

import (
	"context"
	"fmt"
	"time"
	"errors"

	"github.com/google/uuid"
	"github.com/BredSandowich/rss_aggregator/internal/database"
)

func printFeed(name, url, creator string) {
	fmt.Printf("Feed created: %v\n", name)
	fmt.Printf("From: %v\n", url)
	if creator != "" {
		fmt.Printf("Created by: %v\n", creator)
	}
}

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
	
	following, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:      user.ID,
		FeedID: 	feed.ID,
		})
	if err != nil {
		return err
	}

	fmt.Printf("Feed created: %v\n", feed.Name)
	fmt.Printf("From: %v\n", feed.Url)
	fmt.Printf("%s is now following %s\n", following.UserName, following.FeedName)

	//future use print helper function
	//printFeed(feed.Name, feed.Url, user.Name)

	return nil
}


func handlerListFeeds(s *state, cmd command) error {
		//Call method on config pointer
		feeds, err := s.db.GetFeeds(context.Background())
		if err != nil {
			return err
		}

		if len(feeds) == 0 {
			fmt.Println("No feeds to return")
			return nil
		}
	
		for _, feed := range feeds {
			fmt.Printf("Feed created: %v\n", feed.FeedName)
			fmt.Printf("From: %v\n", feed.Url)
			fmt.Printf("Created by: %v\n", feed.UserName)

			//future use print helper function
			//printFeed(feed.FeedName, feed.Url, feed.UserName)
		}

		//Return
		return nil
}