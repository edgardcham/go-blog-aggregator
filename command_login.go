package main

import (
	"context"
	"fmt"
	"os"

	"github.com/edgardcham/go-blog-aggregator/internal/config"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("missing username")
	}

	username := cmd.args[0]

	// check if user exists
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Println("user does not exist")
		os.Exit(1)
	}

	if err := config.SetUser(username, user.ID.String()); err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}

	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	s.cfg = &cfg

	fmt.Printf("Logged in as %s\n", s.cfg.CurrentUserName)

	return nil
}
