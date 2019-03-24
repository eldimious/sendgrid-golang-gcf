package config

import "github.com/eldimious/slack-golang-gcf/utils/env"

// Config is a struct that contains configuration variables
type Config struct {
	SendGrid *SendGrid
}

// Slack is a struct that contains SendGrid's configuration variables
type SendGrid struct {
	APIKey string
}

// NewConfig creates a new Config struct
func NewConfig() (*Config, error) {
	env.CheckDotEnv()
	return &Config{
		SendGrid: &SendGrid{
			APIKey: env.MustGet("API_KEY"),
		},
	}, nil
}
