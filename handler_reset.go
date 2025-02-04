package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	db := s.db
	// CHeck if  user with that name exists
	err := db.ResetUsers(context.Background())
	if err != nil {
		// this means the user already exists
		return fmt.Errorf("Could not reset Users")
	}

	fmt.Println("Users deleted successfully.")

	return nil
}
