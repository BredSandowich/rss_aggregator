package main

import (
	"context"
	"fmt"
	"time"
	"errors"
	"strings"
	"log"

	"github.com/google/uuid"
	"github.com/BredSandowich/rss_aggregator/internal/database"
	"database/sql"
)

func handlerAgg(s *state, cmd command) error {
	
	if len(cmd.Args) != 1 {
		return errors.New("Incorrect arguments entered")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		scrapeFeeds(s)
		}

	return nil
}



func scrapeFeeds(s *state) {
	
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println(err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time: t,
				Valid: true,
			}
		}

		desc := len(item.Description)
		description := sql.NullString{}
		if desc != 0 {
			description = sql.NullString{
				String: item.Description,
				Valid: true,
			}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:  uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID: feed.ID,
			Title: item.Title,
			Description: description,
			Url: item.Link,
			PublishedAt: publishedAt,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
}
