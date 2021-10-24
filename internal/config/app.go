package config

import "os"

// AppConfig is
type AppConfig struct {
	ClientID  string
	Key string
}

// NewAppConfig is
func NewAppConfig() *AppConfig {
	return &AppConfig{
		ClientID:  os.Getenv("OZON_CLIENTID"),
		Key: os.Getenv("OZON_KEY"),
	}
}
