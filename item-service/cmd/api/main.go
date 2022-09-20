package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/vimalkuriensam/item-service/pkg/config"
	"github.com/vimalkuriensam/item-service/pkg/models"
	itemrpc "github.com/vimalkuriensam/item-service/pkg/rpc"
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
	itemCollection := new(itemrpc.ItemCollection)
	if err := rpc.Register(itemCollection); err != nil {
		cfg.Logger.Fatal(err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Env["rpcport"]))
	if err != nil {
		cfg.Logger.Fatal(err)
	}
	cfg.Logger.Printf("RPC server listening on port %v\n", cfg.Env["rpcport"])
	cfg.Logger.Fatal(http.Serve(listener, nil))
}
