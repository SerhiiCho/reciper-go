package apiserver

import "github.com/SerhiiCho/reciper/backend/store"

// Config struct
type Config struct {
	BindAddr string `toml:"bind_addr"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Store:    store.NewConfig(),
	}
}
