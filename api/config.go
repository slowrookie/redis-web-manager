package api

type Config struct {
	Theme    string `json:"theme"`
	Language string `json:"language"`
}

const (
	ConfigC = "config"
	ConfigK = "system_config"
)

var DefaultConfig = &Config{}

func (c *Config) Get() error {
	return GlobalStorage.Read(ConfigC, ConfigK, c)
}

func (c *Config) Set() error {
	return GlobalStorage.Write(ConfigC, ConfigK, *c)
}
