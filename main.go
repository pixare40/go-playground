package main

import (
	"net"
	"os"
	"pixare40/hello/proto/currency"
	"pixare40/hello/server"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(logger)

	// Register the service
	currency.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		logger.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	logger.Info("Starting server on port 9092")
	gs.Serve(l)
}