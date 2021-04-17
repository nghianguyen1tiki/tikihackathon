package config

import "time"

type Server struct {
	Http struct {
		Port string
	}
}

type DB struct {
	DSN     string
	Dialect string
}

type Log struct {
	Mode string
	Type string
}

type Cache struct {
	TTL             time.Duration `mapstructure:"ttl"`
	CleanupInterval time.Duration `mapstructure:"cleanup_interval"`
}

type Feed struct {
	DefaultOffset int `mapstructure:"default_offset"`
	DefaultLimit  int `mapstructure:"default_limit"`
}
