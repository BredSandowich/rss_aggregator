package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	//Call method on config pointer
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for user := 0;  user < len(users); user++ {
		if users[user].Name == s.Config.CurrentUserName {
			fmt.Printf("* %v (current)\n", users[user].Name)
		} else {
			fmt.Printf("* %v\n", users[user].Name)
		}
	}
	//Return
	return nil
}