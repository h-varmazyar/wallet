package transactions

import (
	"context"
	"errors"
	"github.com/google/uuid"
	walletApi "github.com/h-varmazyar/wallet/api/proto"
	"github.com/h-varmazyar/wallet/internal/pkg/db"
	"github.com/h-varmazyar/wallet/internal/pkg/entity"
	"github.com/h-varmazyar/wallet/pkg/mapper"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	walletApi.UnimplementedTransactionServiceServer
	repository Repository
	configs    *Configs
	log        *log.Logger
}

func NewService(configs *Configs, db *db.DB, log *log.Logger) *Service {
	service := &Service{
		repository: NewRepository(db, log),
		configs:    configs,
		log:        log,
	}

	return service
}

func (s *Service) RegisterServer(server *grpc.Server) {
	walletApi.RegisterTransactionServiceServer(server, s)
}

func (s *Service) Create(ctx context.Context, req *walletApi.TransactionCreateReq) (*walletApi.Transaction, error) {
	tx := new(entity.Transaction)
	mapper.Struct(req, tx)

	if tx, err := s.repository.Create(ctx, tx); err != nil {
		return nil, err
	} else {
		response := new(walletApi.Transaction)
		mapper.Struct(tx, response)
		return response, nil
	}
}

func (s *Service) ChangeStatus(ctx context.Context, req *walletApi.TransactionNewStatusReq) (*walletApi.Void, error) {
	var (
		txID uuid.UUID
		err  error
	)

	if txID, err = uuid.Parse(req.ID); err != nil {
		return nil, err
	}

	if !s.canChange(ctx, txID) {
		return nil, errors.New("cannot change transaction status")
	}

	if err = s.repository.ChangeStatus(ctx, txID, req.NewStatus); err != nil {
		return nil, err
	}

	return new(walletApi.Void), nil
}

func (s *Service) List(ctx context.Context, req *walletApi.TransactionListReq) (*walletApi.Transactions, error) {
	var (
		err          error
		walletID     uuid.UUID
		transactions []*entity.Transaction
	)

	if walletID, err = uuid.Parse(req.WalletID); err != nil {
		return nil, err
	}

	if transactions, err = s.repository.List(ctx, walletID); err != nil {
		return nil, err
	}

	response := new(walletApi.Transactions)
	mapper.Slice(transactions, &response.Transactions)

	return response, nil
}

func (s *Service) canChange(ctx context.Context, txID uuid.UUID) bool {
	if tx, err := s.repository.ReturnByID(ctx, txID); err != nil {
		return false
	} else {
		switch tx.Status {
		case walletApi.Transaction_Pending:
			return true
		case walletApi.Transaction_Success, walletApi.Transaction_Failed:
			return false
		}
	}
	return false
}
