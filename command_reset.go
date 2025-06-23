package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.ResetUsers(context.Background()); err != nil {
		return fmt.Errorf("failed to reset users: %w", err)
	}

	fmt.Println("users reset")

	return nil
}
