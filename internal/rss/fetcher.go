package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to send request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch feed: %s", res.Status)
	}

	xmlData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %w", err)
	}

	rssFeed := &RSSFeed{}
	if err := xml.Unmarshal(xmlData, &rssFeed); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal XML: %w", err)
	}
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i := range rssFeed.Channel.Items {
		rssFeed.Channel.Items[i].Title = html.UnescapeString(rssFeed.Channel.Items[i].Title)
		rssFeed.Channel.Items[i].Description = html.UnescapeString(rssFeed.Channel.Items[i].Description)
	}
	return rssFeed, nil
}
