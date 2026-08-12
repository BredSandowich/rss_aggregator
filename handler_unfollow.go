package main

import (
	"context"
	"fmt"
	"errors"

	"github.com/BredSandowich/rss_aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("Incorrect arguments entered")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	
	//Call method on config pointer
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	//Print confirmation of user input
	fmt.Println("Feed has been unfollowed")
	return nil
}