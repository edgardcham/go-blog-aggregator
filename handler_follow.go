package main

import (
	"context"
	"fmt"
	"github.com/edgardcham/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("No feed name provided")
	} else if len(cmd.arguments) > 1 {
		return fmt.Errorf("Too many arguments provided")
	}
	url := cmd.arguments[0]
	db := s.db
	id := uuid.New()
	currentUser := s.config.CURRENT_USER_NAME
	// check if current user exists in db as a safety check
	userData, err := db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("User %s does not exist", currentUser)
	}
	// check if feed exists in db
	feedData, err := db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Feed %s does not exist", url)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userData.ID,
		FeedID:    feedData.ID,
	}
	_, err = db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Failed to create feed follow entry")
	}
	fmt.Printf("Current user: %s is now following feed: %s\n", currentUser, url)

	return nil
}
