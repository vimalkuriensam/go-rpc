package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/vimalkuriensam/broker-service/pkg/config"
	"github.com/vimalkuriensam/broker-service/pkg/routes"
)

const DEFAULT_ENV = "development"

var env string

func main() {
	flag.StringVar(&env, "envflag", DEFAULT_ENV, "Sets the default environment")
	flag.Parse()
	cfg := config.Init()
	cfg.LoadEnvironment(env)
	if err := cfg.ConnectRPC(); err != nil {
		cfg.Logger.Fatalf("could not connect to rpc port %v", cfg.Env["rpcport"])
	}
	routes := routes.Routes()
	cfg.Logger.Printf("Server is running on port %v", cfg.Env["port"])
	cfg.Logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", cfg.Env["port"]), routes))
}
