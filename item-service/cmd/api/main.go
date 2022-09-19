package main

import (
	"flag"

	"github.com/vimalkuriensam/item-service/pkg/config"
)

const DEFAULT_ENVIRONMENT = "development"

var env string

func main() {
	flag.StringVar(&env, "envflag", DEFAULT_ENVIRONMENT, "sets the environment flag")
	flag.Parse()

	config.Initialize().LoadEnvironment(env)
}
