package app

import "WebService/Json"

type Config struct {
	Host         string `json:"HostRedis"`
	Port         string `json:"PortRedis"`
	Password     string `json:"PasswordRedis"`
	DB           int    `json:"DBRedis"`
	SymSplit     string `json:"SYMBOL_SPLIT"`
	HOSTreceiver string `json:"HOSTreceiver"`
	PORTreceiver string `json:"PORTreceiver"`
}

func NewConfig(host, port, pass string, db int, sym, host_rec, port_rec string) *Config {
	conf := Config{}
	conf.Host = host
	conf.Port = port
	conf.Password = pass
	conf.DB = db
	conf.SymSplit = sym
	conf.HOSTreceiver = host_rec
	conf.PORTreceiver = port_rec

	return &conf
}

