package config

import (
	"encoding/json"
	"os"
	"sync"
)

type Config struct {
	Port string `json:"port"`
	DSN  string `json:"dsn"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := instance.configParser("config.json"); err != nil {
			panic(err)
		}
	})
	return instance
}

func (c *Config) configParser(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(c)
}
