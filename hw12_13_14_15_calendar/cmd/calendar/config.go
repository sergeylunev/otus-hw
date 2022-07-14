package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Logger  LoggerConf
	Storage StorageConf
	Server  ServerConf
}

type LoggerConf struct {
	Level     string
	Directory string
	Type      string
}

type StorageConf struct {
	Type string
	User string
	Pass string
	Port int
	Host string
	Name string
}

type ServerConf struct {
	Host string
	Port string
}

func NewConfig(cfg string) Config {
	viper.SetConfigFile(cfg)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	c := Config{}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("cant read confg: %v", err)
	}

	return c
}

func (stc StorageConf) GetDbConnectionString() string {
	return fmt.Sprintf("name=%s dbname=%s password=%s port=%d host=%s sslmode=disable", stc.User, stc.Name, stc.Pass, stc.Port, stc.Host)
}
