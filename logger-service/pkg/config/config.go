package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"github.com/vimalkuriensam/logger-service/pkg/proto"
)

var cfg *Config

type Config struct {
	Env    map[string]any
	Logger *log.Logger
}

type Server struct {
	proto.LogsServiceServer
}

func New() *Config {
	cfg = &Config{
		Env:    make(map[string]any),
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
	return cfg
}

func GetConfig() *Config {
	return cfg
}

func (config *Config) LoadEnvironment(envStatus string) error {
	if envStatus == "production" {
		for _, value := range os.Environ() {
			e := strings.Split(value, "=")
			k, v := e[0], e[1]
			config.Env[k] = v
		}
	} else {
		wd, _ := os.Getwd()
		filePath := filepath.Join(wd, "environment", fmt.Sprintf("%s.env", envStatus))
		viper.SetConfigFile(filePath)
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("error reading config file %v", err)
		}
		for key, value := range viper.AllSettings() {
			config.Env[key] = value
		}
	}
	return nil
}
