package handlers

import (
	"context"
	"fmt"

	"github.com/Sed-Miyuki/RSS_Blog_Aggregator/internal/config"
	"github.com/Sed-Miyuki/RSS_Blog_Aggregator/internal/database"
)
func GetFeedFollowsByUserHandler(s *config.State, cmd config.Command,user database.User) error {
	ctx := context.Background()
	ff, err := s.DB.GetFeedFollowsByUser(ctx, user.ID)
	if err != nil {
		return err
	}
	for _, feed := range ff {
		fmt.Printf("* %s\n", feed.FeedName)
	}
	return nil
}