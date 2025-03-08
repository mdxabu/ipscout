package core

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	WiFi struct {
		Name string `yaml:"name"`
		IPv4 string `yaml:"ipv4"`
		IPv6 string `yaml:"ipv6"`
	} `yaml:"wifi"`
}

func FetchYaml() (string, string, error) {
	filePath := "ipscoutconfig.yaml"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", "", fmt.Errorf("error reading YAML file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return "", "", fmt.Errorf("error parsing YAML: %v", err)
	}

	return config.WiFi.IPv4, config.WiFi.IPv6, nil
}
