package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Env      map[string]any
	Database *Database
	Logger   *log.Logger
	Response *JSONResponse
}

type Database struct {
	Client      *mongo.Client
	Collections map[string]*mongo.Collection
}

type JSONResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var cfg *Config

func Initialize() *Config {
	cfg = &Config{
		Env: map[string]any{},
		Database: &Database{
			Collections: make(map[string]*mongo.Collection),
		},
		Logger:   log.New(os.Stdout, "", log.Ldate|log.Ltime),
		Response: &JSONResponse{},
	}
	return cfg
}

func GetConfig() *Config {
	return cfg
}

func (cfg *Config) LoadEnvironment(envStatus string) error {
	if envStatus == "development" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		path := filepath.Join(wd, "environment", fmt.Sprintf("%s.env", envStatus))
		viper.SetConfigFile(path)
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("error reading config file: %v", err)
		}
		for key, value := range viper.AllSettings() {
			cfg.Env[key] = value
		}
	} else {
		for _, value := range os.Environ() {
			e := strings.Split(value, "=")
			k, v := e[0], e[1]
			cfg.Env[k] = v
		}
	}
	return nil
}
