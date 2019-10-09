package store

// Config struct
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig returns pointer on Config
func NewConfig() *Config {
	return &Config{}
}
