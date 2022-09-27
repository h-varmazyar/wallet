package main

import (
	"github.com/h-varmazyar/wallet/internal/pkg/db"
	"github.com/h-varmazyar/wallet/pkg/netext"
	"github.com/h-varmazyar/wallet/pkg/serverext"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func main() {
	//todo: create new logger
	logger := log.New()
	configs := loadConfigs()
	db := initializeDB(configs.DSN)

	server := serverext.New(logger)
	registerServices(server, configs.GRPCPort)
	registerHandlers(server, configs.HttpPort)
}

func loadConfigs() *Config {
	config := new(Config)
	//todo: load conf
	return config
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

		return grpcServer.Serve(listener)
	})
}

func registerHandlers(server *serverext.Server, port netext.Port) {
	server.Serve(port, func(listener net.Listener) error {

		return http.Serve(listener, handler)
	})
}
