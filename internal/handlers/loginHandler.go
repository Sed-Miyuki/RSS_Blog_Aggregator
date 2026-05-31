package handlers

import (
	"errors"
	"fmt"

	"github.com/Sed-Miyuki/RSS_Blog_Aggregator/internal/config"
)

func HandleLogin(s *config.State,cmd config.Command) error{
	if(len(cmd.Args)<1){
		return errors.New("missing required argument USERNAME")
	}
	err:=s.Config.SetUser(cmd.Args[0])
	if err!=nil{
		return fmt.Errorf("failed to set current user: %w", err)
	}
	fmt.Printf("User set to %s\n", s.Config.CurrentUserName)
	return nil
}