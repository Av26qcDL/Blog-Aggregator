package main

import (
	"context"
	"fmt"
)

func handleFetch(s *state, cmd command) error {
	fmt.Println("Fetch command started") // Add this line
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}

	// Scrape each feed
	for _, feed := range feeds {
		scrapeFeed(s.db, feed) // Pass s.db directly since scrapeFeed expects *database.Queries
	}

	return nil
}
