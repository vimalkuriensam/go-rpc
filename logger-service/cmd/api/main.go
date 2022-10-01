package main

import (
	"flag"

	"github.com/vimalkuriensam/logger-service/pkg/config"
)

const DEFAULT_ENV = "development"

var env string

func main() {
	flag.StringVar(&env, "envflag", DEFAULT_ENV, "Sets the default environment")
	flag.Parse()
	cfg := config.New()
	cfg.LoadEnvironment(env)
	cfg.Logger.Fatal(cfg.RunGRPCServer())
}
