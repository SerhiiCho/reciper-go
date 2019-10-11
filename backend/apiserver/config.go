package apiserver

// Config struct
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
}

// NewConfig returns pointer to Config struct
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
