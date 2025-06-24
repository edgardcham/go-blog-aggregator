package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/edgardcham/go-blog-aggregator/internal/config"
	"github.com/edgardcham/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("missing username")
	}

	username := cmd.args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			// 23505 is the error code for unique constraint violation
			if pgErr.Code == "23505" {
				fmt.Println("username already exists")
				os.Exit(1)
			}
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err = config.SetUser(username, user.ID.String()); err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}

	fmt.Printf("user %s registered\n", user.Name)

	return nil
}
