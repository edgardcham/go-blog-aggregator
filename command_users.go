package main

import (
	"context"
	"fmt"

	"github.com/edgardcham/go-blog-aggregator/internal/config"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}

	// get current user
	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	currentUser := cfg.CurrentUserName

	fmt.Println("users:")
	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("%s (current)\n", user.Name)
		} else {
			fmt.Printf("%s\n", user.Name)
		}
	}

	return nil
}
