package main

import (
	"context"
	"fmt"
	"github.com/edgardcham/go-blog-aggregator/internal/database"
	"github.com/edgardcham/go-blog-aggregator/internal/rss"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Usage: agg <time_between_reqs>")
	}

	duration, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Invalid duration: %w", err)
	}
	fmt.Printf("Collecting feeds every %s\n", duration)

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	// Run immediately, then on every tick.
	scrapeFeeds(s)
	for range ticker.C {
		scrapeFeeds(s)
	}
	return nil
}

func scrapeFeeds(s *state) {
	ctx := context.Background()

	// Get next feed to fetch
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		fmt.Printf("Error getting next feed: %v\n", err)
		return
	}
	if feed.ID == uuid.Nil {
		fmt.Println("No feeds found.")
		return
	}

	// Mark feed as fetched
	if err := s.db.MarkFeedFetched(ctx, feed.ID); err != nil {
		fmt.Printf("Error marking feed fetched: %v\n", err)
		return
	}

	fmt.Printf("Fetching feed: %s\n", feed.Url)
	rssFeed, err := rss.FetchFeed(ctx, feed.Url)
	if err != nil {
		fmt.Printf("Error fetching feed: %v\n", err)
		return
	}

	// Iterate over feed items and save posts to DB
	for _, item := range rssFeed.Channel.Items {
		// Parse published_at; try RFC1123Z then fallback to RFC1123
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			publishedAt, err = time.Parse(time.RFC1123, item.PubDate)
			if err != nil {
				fmt.Printf("Error parsing published date for post %q: %v\n", item.Title, err)
				continue
			}
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		}

		_, err = s.db.CreatePost(ctx, params)
		if err != nil {
			if isDuplicateError(err) {
				continue
			}
			fmt.Printf("Error saving post %q: %v\n", item.Title, err)
		} else {
			fmt.Printf("Saved post: %s\n", item.Title)
		}
	}
}

func isDuplicateError(err error) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		return pgErr.Code == "23505" // unique_violation
	}
	return false
}
