package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var cfg *ini.File
var err error

func init() {
	var envfile string
	if env := os.Getenv("ENV"); env == "development" {
		envfile = ".env.development"
	} else {
		envfile = ".env.production"
	}

	cfg, err = ini.Load(envfile)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}

func GetDBValue(key string) string {
	return GetValue("mysql", key)
}
func GetValue(section string, key string) string {
	return cfg.Section(section).Key(key).String()
}
