package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"`
}

// Returned initial config contains default value
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
