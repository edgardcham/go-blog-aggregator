package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("No username provided")
	} else if len(cmd.arguments) > 1 {
		return fmt.Errorf("Too many arguments provided")
	}
	name := cmd.arguments[0]
	// Check if user exists in the database
	db := s.db
	_, err := db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("User %s does not exist", name)
	}
	if err := s.config.SetUser(name); err != nil {
		return err
	}
	fmt.Println("User set to", name)

	return nil
}
