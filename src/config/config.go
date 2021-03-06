package config

import (
	"encoding/json"
	"io/ioutil"
)

var Config struct {
	Google struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	} `json:"google"`
	Database struct {
		Arguments map[string]string `json:"arguments"`
		DefaultBackend string `json:"default_backend"`
	} `json:"database"`
}

func init() {
	rawConfig, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(rawConfig, &Config); err != nil {
		panic(err)
	}
}
