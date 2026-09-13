// Package config loads the standard-go-web-app configuration from
// environment variables.
package config

import (
	"fmt"

	serverconfig "github.com/tab58/huma-http-server/config"
)

// Config is the app configuration, loaded from environment variables.
type Config struct {
	Port      string `mapstructure:"PORT" default:":8888"`
	JWTSecret string `mapstructure:"JWT_SIGNING_SECRET" sensitive:"true"`
}

// Load reads the configuration from environment variables, logging it (with
// sensitive fields redacted) once loaded.
func Load() (Config, error) {
	var cfg Config
	if err := serverconfig.Load(&cfg, serverconfig.WithConfigDump()); err != nil {
		return Config{}, fmt.Errorf("load config: %w", err)
	}
	if cfg.JWTSecret == "" {
		return Config{}, fmt.Errorf("JWT_SIGNING_SECRET is required")
	}
	return cfg, nil
}
