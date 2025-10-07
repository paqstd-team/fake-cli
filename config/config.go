package config

import (
	"encoding/json"
	"os"
)

type Endpoint struct {
	URL      string      `json:"url"`
	Type     string      `json:"type"`
	Response interface{} `json:"response"`
	Status   int         `json:"status"`
	Payload  interface{} `json:"payload"`
	Cache    *int        `json:"cache"`
}

type Config struct {
	Endpoints []Endpoint `json:"endpoints"`
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
