package main

import "os"
import "io/ioutil"
import "encoding/json"

type Logging struct {
	Categories string `json:"categories"`
	Prefix     string `json:"prefix"`
	Timestamp  bool   `json:"timestamp"`
}

type Listen struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Channel struct {
	ChannelType string `json:"type"`
	Url         string `json:"url"`
}

type Transformation struct {
	TransformationType string `json:"type"`
	Text               string `json:"text"`
	Position           string `json:"position"`
	Anchor             string `json:"anchor"`
}

type Config struct {
	Logging         Logging                   `json:"logging"`
	Listen          Listen                    `json:"listen"`
	Channels        map[string]Channel        `json:"channels"`
	Transformations map[string]Transformation `json:"transformations"`
}

func ReadConfig(name string) (*Config, error) {
	jsonFile, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	json.Unmarshal(byteValue, &config)
	return &config, nil
}
