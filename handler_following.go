package main

import (
	"context"
	"fmt"
	"github.com/edgardcham/go-blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error getting following: %w", err)
	}
	for _, follow := range following {
		fmt.Printf("%s is following %s\n", follow.UserName, follow.FeedName)
	}
	return nil
}
