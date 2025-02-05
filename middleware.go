package main

import (
	"context"
	"github.com/edgardcham/go-blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.CURRENT_USER_NAME)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
