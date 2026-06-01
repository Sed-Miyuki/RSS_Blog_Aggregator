package handlers

import (
	"context"
	"fmt"
	"os"

	rss "github.com/Sed-Miyuki/RSS_Blog_Aggregator/internal/RSS"
	"github.com/Sed-Miyuki/RSS_Blog_Aggregator/internal/config"
)

func FetchFeedHandler(s *config.State, cmd config.Command) error{
	var feedURL string
	if len(cmd.Args)<1{
		feedURL = "https://www.wagslane.dev/index.xml"
	} else {
		feedURL = cmd.Args[0]
	}
	ctx:=context.Background()
	feed, err := rss.FetchFeed(ctx, feedURL)
	if err != nil {
		fmt.Println("Failed to fetch feed")
		os.Exit(1)
	}
	fmt.Printf("%v", feed)
	return nil
}