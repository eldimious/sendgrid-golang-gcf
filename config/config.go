package config

import "github.com/eldimious/sendgrid-golang-gcf/utils/env"

// Config is a struct that contains configuration variables
type Config struct {
	SendGrid *SendGrid
	FaaS     *FaaS
}

// SendGrid is a struct that contains SendGrid's configuration variables
type SendGrid struct {
	APIKey string
}

// SendGrid is a struct that contains SendGrid's configuration variables
type FaaS struct {
	APIKey string
}

// NewConfig creates a new Config struct
func NewConfig() (*Config, error) {
	env.CheckDotEnv()
	return &Config{
		SendGrid: &SendGrid{
			APIKey: env.MustGet("SENDGRID_API_KEY"),
		},
		FaaS: &FaaS{
			APIKey: env.MustGet("API_KEY"),
		},
	}, nil
}
