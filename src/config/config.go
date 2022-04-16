package config

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigStruct struct {
	Google struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	} `json:"google"`
	Database struct {
		Filename string `json:"filename"`
	} `json:"filename"`
}

var Config ConfigStruct

func init() {
	rawConfig, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(rawConfig, &Config); err != nil {
		panic(err)
	}
}
