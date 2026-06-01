package handlers

import (
	"context"
	"fmt"

	"github.com/Sed-Miyuki/RSS_Blog_Aggregator/internal/config"
)

func ListFeedHandler(s *config.State,cmd config.Command) error{
	ctx := context.Background()
	feeds, err := s.DB.ListFeeds(ctx)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("Article: %s | Url: %s | User: %s\n", feed.Name, feed.Url, feed.UserName)
	}
	return nil
}