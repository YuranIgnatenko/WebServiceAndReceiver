package receiver

import (
	Json "WebReceiver/Json"
)

type Config struct {
	Host        string `json:"HOST"`
	Port        string `json:"PORT"`
	SymSplit    string `json:"SYMBOL_SPLIT"`
	TypeConnect string `json:"TYPE_CONNECT"`
	PathLog     string `json:"PATH_LOG"`
}

func NewConfig(filepath string) *Config {
	conf := Config{}
	err := Json.JsonToStruct(filepath, &conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
