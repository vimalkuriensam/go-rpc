package main

import (
	"flag"

	"github.com/vimalkuriensam/item-service/pkg/config"
	"github.com/vimalkuriensam/item-service/pkg/models"
)

const DEFAULT_ENVIRONMENT = "development"

var env string

func main() {
	// loads command line string variable envflag with
	// default environment if not present
	flag.StringVar(&env, "envflag", DEFAULT_ENVIRONMENT, "sets the environment flag")
	// Parses all the flag args
	flag.Parse()
	// Initializes the environment and stores into config
	cfg := config.Initialize()
	if err := cfg.LoadEnvironment(env); err != nil {
		cfg.Logger.Fatal(err)
	}
	db := models.Init()
	if err := db.Connect(); err != nil {
		cfg.Logger.Fatal(err)
	}
	defer db.Disconnect()
	db.InsertMongoCollections("items")
}
