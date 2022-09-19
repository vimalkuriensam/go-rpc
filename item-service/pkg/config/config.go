package config

import "go.mongodb.org/mongo-driver/mongo"

type Config struct {
	Env      map[string]any
	Database *Database
}

type Database struct {
	Collections map[string]*mongo.Collection
}

var cfg *Config

func Initialize() *Config {
	cfg = &Config{
		Env: map[string]any{},
		Database: &Database{
			Collections: make(map[string]*mongo.Collection),
		},
	}
	return cfg
}

func GetConfig() *Config {
	return cfg
}

func LoadEnvironment(envStatus string) error {
	return nil
}
