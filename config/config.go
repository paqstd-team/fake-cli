package config

import (
	"encoding/json"
	"os"
)

type Endpoint struct {
	URL      string `json:"url"`
	Fields   any    `json:"fields"`
	Response string `json:"response"`
}

type Config struct {
	Endpoints []Endpoint `json:"endpoints"`
	Cache     int        `json:"cache"`
}

func LoadConfigFromFile(path string) (Config, error) {
	var config Config

	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
