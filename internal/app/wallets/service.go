package wallets

import (
	"context"
	"errors"
	walletApi "github.com/h-varmazyar/wallet/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	walletApi.UnimplementedWalletServiceServer
	configs *ServiceConfigs
	log     *log.Logger
}

func NewService(configs *ServiceConfigs, log *log.Logger) *Service {
	service := &Service{
		configs: configs,
		log:     log,
	}
	return service
}

func (s *Service) RegisterServer(server *grpc.Server) {
	walletApi.RegisterWalletServiceServer(server, s)
}

func (s *Service) Deposit(ctx context.Context, req *walletApi.NewTransaction) (*walletApi.Wallet, error) {
	return new(walletApi.Wallet), errors.New("UnImplemented")
}

func (s *Service) Withdrawal(ctx context.Context, req *walletApi.NewTransaction) (*walletApi.Wallet, error) {
	return new(walletApi.Wallet), errors.New("UnImplemented")
}

func (s *Service) ReturnByPhoneNumber(ctx context.Context, req *walletApi.WalletReturnByPhoneReq) (*walletApi.Wallet, error) {
	return new(walletApi.Wallet), errors.New("UnImplemented")
}
