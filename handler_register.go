package main

import (
	"context"
	"fmt"
	"time"

	"github.com/edgardcham/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("No name provided")
	} else if len(cmd.arguments) > 1 {
		return fmt.Errorf("Too many arguments provided")
	}
	name := cmd.arguments[0]
	db := s.db
	// CHeck if  user with that name exists
	_, err := db.GetUser(context.Background(), name)
	if err == nil {
		// this means the user already exists
		return fmt.Errorf("User %s already exists", name)
	}
	id := uuid.New()
	params := database.CreateUserParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	user, err := db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error creating user: %w", err)
	}
	// Set the user in the config to the given name
	if err := s.config.SetUser(user.Name); err != nil {
		return fmt.Errorf("Error setting user: %w", err)
	}
	fmt.Println(fmt.Sprintf("User created, user params:\n\v", user))

	return nil
}
