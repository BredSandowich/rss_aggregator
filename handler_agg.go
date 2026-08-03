package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	agg, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(agg)
	return nil
}
