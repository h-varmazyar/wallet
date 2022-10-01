package transactions

import (
	"context"
	"github.com/google/uuid"
	walletApi "github.com/h-varmazyar/wallet/api/proto"
	"github.com/h-varmazyar/wallet/internal/pkg/db"
	"github.com/h-varmazyar/wallet/internal/pkg/entity"
	"github.com/h-varmazyar/wallet/pkg/validatorext"
	log "github.com/sirupsen/logrus"
)

type Repository interface {
	Create(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
	ReturnByID(ctx context.Context, txID uuid.UUID) (*entity.Transaction, error)
	ChangeStatus(ctx context.Context, txID uuid.UUID, newStatus walletApi.TransactionStatus) error
	List(ctx context.Context, walletID uuid.UUID) ([]*entity.Transaction, error)
}

type repository struct {
	logger *log.Logger
	db     *db.DB
}

func NewRepository(db *db.DB, logger *log.Logger) Repository {
	return &repository{logger, db}
}

func (r repository) Create(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	if err := validatorext.Struct(ctx, transaction); err != nil {
		return nil, err
	}
	if err := r.db.Save(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r repository) ReturnByID(_ context.Context, txID uuid.UUID) (*entity.Transaction, error) {
	transaction := new(entity.Transaction)
	if err := r.db.
		Model(new(entity.Transaction)).
		Where("id = ?", txID).
		First(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r repository) ChangeStatus(_ context.Context, txID uuid.UUID, newStatus walletApi.TransactionStatus) error {
	if err := r.db.
		Model(new(entity.Transaction)).
		Where("id = ?", txID).
		Update("status", newStatus).Error; err != nil {
		return err
	}
	return nil
}

func (r repository) List(_ context.Context, walletID uuid.UUID) ([]*entity.Transaction, error) {
	transactions := make([]*entity.Transaction, 0)
	if err := r.db.
		Model(new(entity.Transaction)).
		Where("wallet_id = ?", walletID).
		Find(transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
