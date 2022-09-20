package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Env           map[string]any
	Logger        *log.Logger
	Response      *JSONResponse
	ErrorResponse *ErrorResponse
}

type JSONResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Code      int       `json:"errorCode"`
	Path      string    `json:"path"`
	Reason    string    `json:"reasone"`
	Timestamp time.Time `json:"timeStamp"`
}

var cfg *Config

func Init() *Config {
	cfg = &Config{
		Env:           make(map[string]any),
		Logger:        log.New(os.Stdout, "", log.Ldate|log.Ltime),
		Response:      &JSONResponse{},
		ErrorResponse: &ErrorResponse{},
	}
	return cfg
}

func GetConfig() *Config {
	return cfg
}
