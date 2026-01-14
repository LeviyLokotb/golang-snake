package config

import (
	"encoding/json"
	"os"
)

func LoadConfigFromJSON(path string) (*GameConfig, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	conf := NewDefaultConfig()
	err = json.Unmarshal(content, &conf)

	if err != nil {
		return nil, err
	}

	return conf, nil
}
