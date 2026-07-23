package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	//Call method on config pointer
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	//Print confirmation of user input
	fmt.Println("Users have been deleted")
	return nil
}