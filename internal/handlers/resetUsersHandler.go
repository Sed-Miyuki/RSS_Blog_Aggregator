package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/Sed-Miyuki/RSS_Blog_Aggregator/internal/config"
)
func ResetUsersHandler(s *config.State, cmd config.Command) error {
	ctx := context.Background()
	err := s.DB.ResetUsers(ctx)
	if err != nil {
		fmt.Println("Failed to remove users")
		os.Exit(1)
	}
	fmt.Println("Users removed")
	return nil
}