package wallets

import (
	"context"
	"errors"
	"github.com/google/uuid"
	walletApi "github.com/h-varmazyar/wallet/api/proto"
	"github.com/h-varmazyar/wallet/internal/pkg/db"
	"github.com/h-varmazyar/wallet/internal/pkg/entity"
	"github.com/h-varmazyar/wallet/pkg/grpcext"
	"github.com/h-varmazyar/wallet/pkg/mapper"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	walletApi.UnimplementedWalletServiceServer
	repository         Repository
	configs            *Configs
	log                *log.Logger
	transactionService walletApi.TransactionServiceClient
}

func NewService(configs *Configs, db *db.DB, logger *log.Logger) *Service {
	walletConn := grpcext.NewConnection(configs.TransactionServiceAddress)
	service := &Service{
		repository:         NewRepository(db, logger),
		configs:            configs,
		log:                logger,
		transactionService: walletApi.NewTransactionServiceClient(walletConn),
	}
	return service
}

func (s *Service) RegisterServer(server *grpc.Server) {
	walletApi.RegisterWalletServiceServer(server, s)
}

func (s *Service) Deposit(ctx context.Context, req *walletApi.NewTransaction) (*walletApi.Wallet, error) {
	var (
		err             error
		walletID        uuid.UUID
		wallet          *entity.Wallet
		walletAvailable bool
		tx              *walletApi.Transaction
	)
	if req.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	if walletID, err = uuid.Parse(req.WalletID); err != nil {
		return nil, err
	}

	if walletAvailable, err = s.repository.IsAvailable(ctx, walletID); err != nil || !walletAvailable {
		return nil, err
	}

	if tx, err = s.transactionService.Create(ctx, &walletApi.TransactionCreateReq{
		Type:        walletApi.Transaction_Deposit,
		Status:      walletApi.Transaction_Pending,
		Amount:      req.Amount,
		WalletID:    req.WalletID,
		Description: req.Description,
		RefID:       req.RefID,
	}); err != nil {
		return nil, err
	}

	if err = s.repository.Deposit(ctx, walletID, req.Amount); err != nil {
		if err = s.changeTxStatus(ctx, tx.ID, walletApi.Transaction_Failed); err != nil {
			//todo: log error
		}
		return nil, errors.New("failed to deposit wallet")
	}

	if err = s.changeTxStatus(ctx, tx.ID, walletApi.Transaction_Success); err != nil {
		if err = s.repository.Withdrawal(ctx, walletID, req.Amount); err != nil {
			//todo: log error
		}
		return nil, errors.New("failed to deposit wallet")
	}

	if wallet, err = s.repository.ReturnByID(ctx, walletID); err != nil {
		return nil, err
	}

	response := new(walletApi.Wallet)
	mapper.Struct(wallet, response)
	return response, nil
}

func (s *Service) Withdrawal(ctx context.Context, req *walletApi.NewTransaction) (*walletApi.Wallet, error) {
	var (
		err             error
		walletID        uuid.UUID
		wallet          *entity.Wallet
		walletAvailable bool
		tx              *walletApi.Transaction
	)
	if req.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	if walletID, err = uuid.Parse(req.WalletID); err != nil {
		return nil, err
	}

	if walletAvailable, err = s.repository.IsAvailable(ctx, walletID); err != nil || !walletAvailable {
		return nil, err
	}

	if tx, err = s.transactionService.Create(ctx, &walletApi.TransactionCreateReq{
		Type:        walletApi.Transaction_Withdraw,
		Status:      walletApi.Transaction_Pending,
		Amount:      req.Amount,
		WalletID:    req.WalletID,
		Description: req.Description,
		RefID:       req.RefID,
	}); err != nil {
		return nil, err
	}

	if err = s.repository.Withdrawal(ctx, walletID, req.Amount); err != nil {
		if err = s.changeTxStatus(ctx, tx.ID, walletApi.Transaction_Failed); err != nil {
			//todo: log error
		}
		return nil, errors.New("failed to deposit wallet")
	}

	if err = s.changeTxStatus(ctx, tx.ID, walletApi.Transaction_Success); err != nil {
		if err = s.repository.Deposit(ctx, walletID, req.Amount); err != nil {
			//todo: log error
		}
		return nil, errors.New("failed to deposit wallet")
	}

	if wallet, err = s.repository.ReturnByID(ctx, walletID); err != nil {
		return nil, err
	}

	response := new(walletApi.Wallet)
	mapper.Struct(wallet, response)
	return response, nil
}

func (s *Service) ReturnByPhoneNumber(ctx context.Context, req *walletApi.WalletReturnByPhoneReq) (*walletApi.Wallet, error) {
	//todo: validate phone number
	if wallet, err := s.repository.ReturnByPhoneNumber(ctx, req.PhoneNumber); err != nil {
		return nil, err
	} else {
		response := new(walletApi.Wallet)
		mapper.Struct(wallet, response)
		return response, nil
	}
}

func (s *Service) changeTxStatus(ctx context.Context, txID string, newStatus walletApi.TransactionStatus) error {
	if _, err := s.transactionService.ChangeStatus(ctx, &walletApi.TransactionNewStatusReq{
		ID:        txID,
		NewStatus: newStatus,
	}); err != nil {
		return err
	}
	return nil
}
