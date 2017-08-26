package main

import (
	//"github.com/BurntSushi/toml"
)

type DatabaseConfig struct {
	// Postgres database config
	IP string
	Port uint16
	Username string
	Password string
	Database string
}

type ServerConfig struct {
	IPAddress string `toml:"ip"`
	Port uint16 `toml:"port"`
	MaxPlayer uint32 `toml:"max_player"`
	Database DatabaseConfig

}

type Config struct {
	Server ServerConfig `toml:"server"`

}

func ParseArgument() {
	//toml.Decode()
}

