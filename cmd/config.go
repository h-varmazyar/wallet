package main

import "github.com/h-varmazyar/wallet/pkg/netext"

type Config struct {
	DSN      string      `yaml:"dsn"`
	GRPCPort netext.Port `yaml:"grpc_port"`
	HttpPort netext.Port `yaml:"http_port"`
}
