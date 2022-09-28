package main

import (
	"github.com/h-varmazyar/wallet/internal/app/transactions"
	"github.com/h-varmazyar/wallet/internal/app/wallets"
	"github.com/h-varmazyar/wallet/internal/pkg/db"
	"github.com/h-varmazyar/wallet/pkg/netext"
	"github.com/h-varmazyar/wallet/pkg/serverext"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

var (
	logger  *log.Logger
	configs *Configs
)

const (
	name    = "Wallet"
	version = "v1.0.0"
)

func main() {
	//todo: create new logger
	logger = log.New()
	configs = loadConfigs()
	db := initializeDB(configs.DSN)

	server := serverext.New(logger)
	registerServices(server, configs.GRPCPort)
	registerHandlers(server, configs.HttpPort)

	server.Start(name, version)
}

func loadConfigs() *Configs {
	configs := new(Configs)
	//todo: load conf
	return configs
}

func initializeDB(dsn string) *db.DB {
	dbInstance, err := db.New(dsn)
	if err != nil {
		log.Panicf(`failed to instantiate new database: %v`, err)
	}
	return dbInstance
}

func registerServices(server *serverext.Server, port netext.Port) {
	server.Serve(port, func(listener net.Listener) error {
		grpcServer := grpc.NewServer()

		transactions.NewService(configs.TransactionsConfigs, logger).RegisterServer(grpcServer)
		wallets.NewService(configs.WalletsConfigs, logger).RegisterServer(grpcServer)

		return grpcServer.Serve(listener)
	})
}

func registerHandlers(server *serverext.Server, port netext.Port) {
	server.Serve(port, func(listener net.Listener) error {

		return http.Serve(listener, handler)
	})
}
