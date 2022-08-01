package main

import (
	"log"

	"github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/spf13/viper"
)

type Config struct {
	Logger  LoggerConf
	Storage storage.StorageConf
	Server  HttpConf
	Grpc    GrpcConf
}

type LoggerConf struct {
	Level     string
	Directory string
	Type      string
}

type HttpConf struct {
	Host string
	Port string
}

type GrpcConf struct {
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
