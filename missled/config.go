package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

const (
	DSN = "root:apple@tcp(localhost:3306)/auth?charset=utf8"
	MC  = "localhost:11211"
)

type Config struct {
	MysqlDsn    string
	MemcacheDsn string
}

func loadConfig() (config Config, err error) {
	if _, err = toml.DecodeFile("config.toml", &config); err != nil {
		log.Printf("read config file error: %s\n", err.Error())
		return
	}

	return
}
