package config

import "log"

type ConfigServer struct {
	Port       int
	StaticPath string
}

type Application struct {
	ConfigServer *ConfigServer
	Logger       *log.Logger
}
