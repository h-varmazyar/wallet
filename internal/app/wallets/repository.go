package wallets

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/h-varmazyar/wallet/internal/pkg/db"
	"github.com/h-varmazyar/wallet/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, wallet *entity.Wallet) (*entity.Wallet, error)
	IsAvailable(ctx context.Context, walletID uuid.UUID) (bool, error)
	ReturnByID(ctx context.Context, walletID uuid.UUID) (*entity.Wallet, error)
	ReturnByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.Wallet, error)
	Deposit(ctx context.Context, walletID uuid.UUID, amount int64) error
	Withdrawal(ctx context.Context, walletID uuid.UUID, amount int64) error
}

type repository struct {
	logger *log.Logger
	db     *db.DB
}

func NewRepository(db *db.DB, logger *log.Logger) Repository {
	return &repository{logger, db}
}

func (r *repository) Create(_ context.Context, wallet *entity.Wallet) (*entity.Wallet, error) {
	//TODO validate wallet
	count := int64(0)
	if err := r.db.
		Model(new(entity.Wallet)).
		Where("phone_number = ?", wallet.PhoneNumber).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("wallet available")
	}

	if err := r.db.Save(wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *repository) IsAvailable(_ context.Context, id uuid.UUID) (bool, error) {
	//todo: validate wallet
	count := int64(0)
	if err := r.db.
		Model(new(entity.Wallet)).
		Where("id = ?", id).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repository) ReturnByID(_ context.Context, id uuid.UUID) (*entity.Wallet, error) {
	//todo: validate wallet
	wallet := new(entity.Wallet)
	if err := r.db.
		Model(new(entity.Wallet)).
		Where("id = ?", id).
		First(wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *repository) ReturnByPhoneNumber(_ context.Context, phoneNumber string) (*entity.Wallet, error) {
	//todo: validate wallet
	wallet := new(entity.Wallet)
	if err := r.db.
		Model(new(entity.Wallet)).
		Where("phone_number = ?", phoneNumber).
		First(wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *repository) Deposit(_ context.Context, walletID uuid.UUID, amount int64) error {
	//todo: save tx there or not
	if err := r.db.
		Model(new(entity.Wallet)).
		Where("id = ?", walletID).
		Update("amount", gorm.Expr("amount + ?", amount)).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Withdrawal(_ context.Context, walletID uuid.UUID, amount int64) error {
	//todo: save tx there or not
	wallet := new(entity.Wallet)
	if err := r.db.
		Where("id = ?", walletID).
		First(wallet).Error; err != nil {
		return err
	}
	if wallet.Amount < amount {
		return errors.New("wallet amount insufficient, cannot withdrawal with this amount")
	}
	if err := r.db.
		Model(new(entity.Wallet)).
		Where("id = ?", walletID).
		Update("amount", gorm.Expr("amount - ?", amount)).Error; err != nil {
		return err
	}
	return nil
}
