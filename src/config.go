package main

import "os"
import "io/ioutil"
import "encoding/json"

type loggingConfig struct {
	Categories string `json:"categories"`
	Prefix     string `json:"prefix"`
	Timestamp  bool   `json:"timestamp"`
}

type listenConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type channelConfig struct {
	ChannelType string `json:"type"`
	Url         string `json:"url"`
	Render      string `json:"render"`
}

type transformationConfig struct {
	TransformationType string            `json:"type"`
	Params             map[string]string `json:"params"`
}

type config struct {
	Logging         loggingConfig                   `json:"logging"`
	Listen          listenConfig                    `json:"listen"`
	Channels        map[string]channelConfig        `json:"channels"`
	Transformations map[string]transformationConfig `json:"transformations"`
}

func ReadConfig(name string) (*config, error) {
	jsonFile, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var cfg config
	err = json.Unmarshal(byteValue, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
