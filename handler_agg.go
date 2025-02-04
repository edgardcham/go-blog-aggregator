package main

import (
	"context"
	"fmt"
	"github.com/edgardcham/go-blog-aggregator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	rssFeed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %w", err)
	}
	fmt.Printf("Feed fetched: %s\n", rssFeed)
	return nil
}
