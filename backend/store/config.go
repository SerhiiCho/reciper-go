package store

type Config struct {
	DBPwd  string `toml:"db_pwd"`
	DBUser string `toml:"db_user"`
	DBHost string `toml:"db_host"`
	DBPort string `toml:"db_port"`
	DBName string `toml:"db_name"`
}

func NewConfig() *Config {
	return &Config{
		DBPwd:  "111111",
		DBUser: "root",
		DBHost: "localhost",
		DBPort: "1001",
		DBName: "reciper",
	}
}
