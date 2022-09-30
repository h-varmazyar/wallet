package main

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/wallet/internal/app/transactions"
	"github.com/h-varmazyar/wallet/internal/app/wallets"
	"github.com/h-varmazyar/wallet/internal/pkg/db"
	"github.com/h-varmazyar/wallet/internal/pkg/entity"
	"github.com/h-varmazyar/wallet/pkg/netext"
	"github.com/h-varmazyar/wallet/pkg/serverext"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"io/ioutil"
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
	logger = log.StandardLogger()
	configs = loadConfigs()

	dbInstance := initializeDB(configs.DSN)

	server := serverext.New(logger)
	registerServices(server, dbInstance, configs.GRPCPort)
	registerHandlers(server, configs.HttpPort)

	server.Start(name, version)
}

func loadConfigs() *Configs {
	configs := new(Configs)

	confBytes, err := ioutil.ReadFile("configs/default.yaml")
	if err != nil {
		log.WithError(err).Fatal("can not load yaml file")
	}
	if err = yaml.Unmarshal(confBytes, configs); err != nil {
		log.WithError(err).Fatal("can not unmarshal yaml file")
	}

	return configs
}

func initializeDB(dsn string) *db.DB {
	dbInstance, err := db.New(dsn)
	if err != nil {
		log.Panicf(`failed to instantiate new database: %v`, err)
	}
	if err = dbInstance.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
			return err
		}
		if err = migrateModels(tx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Panicf(`failed to migrate entities: %v`, err)
	}

	return dbInstance
}

func migrateModels(db *gorm.DB) error {
	if err := db.AutoMigrate(
		new(entity.Wallet),
		new(entity.Transaction),
	); err != nil {
		return err
	}
	return nil
}

func registerServices(server *serverext.Server, dbInstance *db.DB, port netext.Port) {
	server.Serve(port, func(listener net.Listener) error {
		grpcServer := grpc.NewServer()

		transactions.NewService(configs.TransactionsConfigs, dbInstance, logger).RegisterServer(grpcServer)
		wallets.NewService(configs.WalletsConfigs, dbInstance, logger).RegisterServer(grpcServer)

		return grpcServer.Serve(listener)
	})
}

func registerHandlers(server *serverext.Server, port netext.Port) {
	server.Serve(port, func(listener net.Listener) error {
		router := gin.Default()

		wallets.NewHandler(configs.WalletsConfigs, logger).RegisterRoutes(router)

		return http.Serve(listener, router)
	})
}
