package main

import (
	"github.com/h-varmazyar/wallet/internal/app/transactions"
	"github.com/h-varmazyar/wallet/internal/app/wallets"
	"github.com/h-varmazyar/wallet/pkg/netext"
)

type Configs struct {
	DSN                 string                `yaml:"dsn"`
	GRPCPort            netext.Port           `yaml:"grpc_port"`
	HttpPort            netext.Port           `yaml:"http_port"`
	TransactionsConfigs *transactions.Configs `yaml:"transactions_configs"`
	WalletsConfigs      *wallets.Configs      `yaml:"wallets_configs"`
}
