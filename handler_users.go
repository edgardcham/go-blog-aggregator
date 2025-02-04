package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {

	db := s.db
	users, err := db.GetUsers(context.Background())
	if err != nil {
		// this means the user already exists
		return fmt.Errorf("Could not get Users")
	}
	currentUser := s.config.CURRENT_USER_NAME

	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
