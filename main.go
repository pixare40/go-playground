package main

import (
	"pixare40/hello/proto/currency"
	"pixare40/hello/server"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	logger := hclog.Default()

	gs := grpc.NewServer()

	// Register the service
	currency.RegisterCurrencyServer(gs, server.NewCurrency(logger))
}