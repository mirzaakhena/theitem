package config

import (
	"encoding/json"
	"os"
)

const defaultConfigFile = "config.json"

func ReadConfig() *Config {

	configFile := defaultConfigFile
	configFileFromEnv := os.Getenv("CONFIG_FILE")
	if configFileFromEnv != "" {
		configFile = configFileFromEnv
	}

	bytes, err := os.ReadFile(configFile)
	if err != nil {
		panic(err.Error())
	}

	var cfg Config

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		panic(err.Error())
	}

	return &cfg
}
