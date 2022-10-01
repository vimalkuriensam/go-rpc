package config

import (
	"net"

	"github.com/vimalkuriensam/logger-service/pkg/proto"
	"google.golang.org/grpc"
)

func (cfg *Config) RunGRPCServer() error {
	listener, err := net.Listen("tcp", cfg.Env["rpcport"].(string))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	proto.RegisterLogsServiceServer(s, &Server{})
	cfg.Logger.Printf("Listening on %s\n", listener.Addr())
	return s.Serve(listener)
}
