package transactions

import (
	"context"
	"errors"
	walletApi "github.com/h-varmazyar/wallet/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	walletApi.UnimplementedTransactionServiceServer
	configs *Configs
	log     *log.Logger
}

func NewService(configs *Configs, log *log.Logger) *Service {
	service := &Service{
		configs: configs,
		log:     log,
	}

	return service
}

func (s Service) RegisterServer(server *grpc.Server) {
	walletApi.RegisterTransactionServiceServer(server, s)
}

func (s Service) List(ctx context.Context, req *walletApi.TransactionListReq) (*walletApi.Transactions, error) {
	return new(walletApi.Transactions), errors.New("UnImplemented")
}
