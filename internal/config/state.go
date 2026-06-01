package config

import "github.com/Sed-Miyuki/RSS_Blog_Aggregator/internal/database"

type State struct{
	DB     *database.Queries
	Config *Config
}