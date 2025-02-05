package main

import (
	"context"
	"fmt"
	"time"

	"github.com/edgardcham/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("Missing parameters. Usage: addfeed <name> <feedURL>")
	}

	name := cmd.arguments[0]
	feedURL := cmd.arguments[1]
	userId := user.ID // Use the user provided by middleware

	id := uuid.New()
	parameters := database.CreateFeedParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       feedURL,
		UserID:    userId,
	}

	feed, err := s.db.CreateFeed(context.Background(), parameters)
	if err != nil {
		return fmt.Errorf("Error creating feed: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userId,
		FeedID:    feed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Failed to create feed follow entry: %w", err)
	}

	fmt.Println(fmt.Sprintf("Feed created, feed params:\n%v", feed))
	return nil
}
