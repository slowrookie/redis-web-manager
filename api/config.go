package api

type Config struct {
	Theme    string `json:"theme"`
	Language string `json:"language"`
}

const (
	configStorageCollection = "config"
	configStorageKey        = "system_config"
)

var DefaultConfig = &Config{}

func (c *Config) Get() error {
	return GlobalStorage.Read(configStorageCollection, configStorageKey, c)
}

func (c *Config) Set() error {
	return GlobalStorage.Write(configStorageCollection, configStorageKey, *c)
}
