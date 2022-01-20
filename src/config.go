package main

import "os"
import "io/ioutil"
import "encoding/json"

type LoggingConfig struct {
	Categories string `json:"categories"`
	Prefix     string `json:"prefix"`
	Timestamp  bool   `json:"timestamp"`
}

type ListenConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ChannelConfig struct {
	ChannelType string `json:"type"`
	Url         string `json:"url"`
}

type TransformationConfig struct {
	TransformationType string `json:"type"`
	Text               string `json:"text"`
	Position           string `json:"position"`
	Anchor             string `json:"anchor"`
}

type Config struct {
	Logging         LoggingConfig                   `json:"logging"`
	Listen          ListenConfig                    `json:"listen"`
	Channels        map[string]ChannelConfig        `json:"channels"`
	Transformations map[string]TransformationConfig `json:"transformations"`
}

func ReadConfig(name string) (*Config, error) {
	jsonFile, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
