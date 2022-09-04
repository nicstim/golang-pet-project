package apiserver

type Config struct {
	address  string `toml:"bind_address"`
	logLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{address: ":4000", logLevel: "debug"}
}
