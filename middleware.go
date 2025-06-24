package main

import (
	"context"
	"fmt"

	"github.com/edgardcham/go-blog-aggregator/internal/config"
	"github.com/edgardcham/go-blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		cfg, err := config.Read()
		if err != nil {
			return fmt.Errorf("failed to read config: %w", err)
		}

		user, err := s.db.GetUser(context.Background(), cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}

		return handler(s, cmd, user)
	}
}