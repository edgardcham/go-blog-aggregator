package main

import (
	"context"

	"fmt"

	"github.com/edgardcham/go-blog-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("No url provided")
	} else if len(cmd.arguments) > 1 {
		return fmt.Errorf("Too many arguments provided")
	}
	feedUrl := cmd.arguments[0]

	deleteFeedFollowParams := database.DeleteFeedFollowForUserParams{
		Name: user.Name,
		Url:  feedUrl,
	}

	db := s.db
	err := db.DeleteFeedFollowForUser(context.Background(), deleteFeedFollowParams)
	if err != nil {
		return fmt.Errorf("Could not delete feed follow for %s", user.Name)
	}

	fmt.Println("Feed follow deleted for %s", user.Name)

	return nil
}
