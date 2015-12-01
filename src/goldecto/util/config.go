package util

import (
	"io/ioutil"
	
	"gopkg.in/yaml.v2"
)

type GoldEctoConfig struct {
	Db struct {
		Driver string
		Source string
	}
	Web struct {
		Host string
		Port int
	}
}

func ConfigFromData(data []byte) (conf GoldEctoConfig) {
	err := yaml.Unmarshal(data, &conf)
	if (err != nil) { panic(err) }
	return
}

func ConfigFromFile(filename string) GoldEctoConfig {
	data, err := ioutil.ReadFile(filename)
	if (err != nil) { panic(err) }
	return ConfigFromData(data)
}
