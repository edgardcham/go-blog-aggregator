package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	currentUser := s.config.CURRENT_USER_NAME

	// fetch current user data
	userData, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("Error getting user: %w", err)
	}
	userId := userData.ID
	db := s.db
	following, err := db.GetFeedFollowsForUser(context.Background(), userId)
	if err != nil {
		return fmt.Errorf("Error getting following: %w", err)
	}
	for _, follow := range following {
		fmt.Printf("%s is following %s\n", follow.UserName, follow.FeedName)
	}
	return nil
}
