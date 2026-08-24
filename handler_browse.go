package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BredSandowich/rss_aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		limit = parsedLimit
	}
	
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return err
}	
	
	for _, post := range posts {
		fmt.Printf("%s\n%s\n\n", post.Title, post.Url)
	}

	return nil
}