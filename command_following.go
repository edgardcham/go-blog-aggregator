package main

import (
	"context"
	"fmt"

	"github.com/edgardcham/go-blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("failed to get feed follows: %w", err)
	}

	for _, feedFollow := range feedFollows {
		fmt.Printf("Feed Name: %s\n", feedFollow.FeedName)
		fmt.Printf("Feed URL: %s\n", feedFollow.FeedUrl)
	}

	return nil
}
