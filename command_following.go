package main

import (
	"context"
	"fmt"

	"github.com/edgardcham/go-blog-aggregator/internal/config"
)

func handlerFollowing(s *state, cmd command) error {
	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get feed follows: %w", err)
	}

	for _, feedFollow := range feedFollows {
		fmt.Printf("Feed Name: %s\n", feedFollow.FeedName)
		fmt.Printf("Feed URL: %s\n", feedFollow.FeedUrl)
	}

	return nil
}
