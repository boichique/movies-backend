package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	DBUrl      string           `env:"DB_URL"`
	Port       int              `env:"PORT" envDefault:"8080"`
	Local      bool             `env:"LOCAL" envDefault:"false"`
	LogLevel   string           `env:"LOG_LEVEL" envDefault:"info"`
	Jwt        JwtConfig        `envPrefix:"JWT_"`
	Admin      AdminConfig      `envPrefix:"ADMIN_"`
	Pagination PaginationConfig `envPrefix:"PAGINATION_"`
}

type JwtConfig struct {
	Secret           string        `env:"SECRET"`
	AccessExpiration time.Duration `env:"ACCESS_EXPIRATION" envDefault:"15m"`
}

type AdminConfig struct {
	AdminName     string `env:"NAME" validate:"min=5,max=16"`
	AdminEmail    string `env:"EMAIL" validate:"email"`
	AdminPassword string `env:"PASSWORD" validate:"password"`
}

type PaginationConfig struct {
	DefaultSize int `env:"DEFAULT_SIZE" envDefault:"10"`
	MaxSize     int `env:"MAX_SIZE" envDefault:"100"`
}

func NewConfig() (*Config, error) {
	var c Config
	if err := env.Parse(&c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &c, nil
}

func (ac *AdminConfig) AdminIsSet() bool {
	return ac.AdminName != "" && ac.AdminEmail != "" && ac.AdminPassword != ""
}
