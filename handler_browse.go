package main

import (
	"context"
	"fmt"
	"github.com/edgardcham/go-blog-aggregator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.arguments) > 0 {
		limitInt, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("Invalid limit parameter: %w", err)
		}
		limit = int32(limitInt)
	}

	params := database.GetPostsForUserParams{
		Name:  user.Name,
		Limit: limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error fetching posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found.")
		return nil
	}

	for _, p := range posts {
		fmt.Printf("Title: %s\nURL: %s\nPublished: %v\n\n", p.Title, p.Url, p.PublishedAt)
	}
	return nil
}
