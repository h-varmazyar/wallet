package entity

import (
	"github.com/google/uuid"
	walletApi "github.com/h-varmazyar/wallet/api/proto"
	"github.com/h-varmazyar/wallet/internal/pkg/db"
)

type Transaction struct {
	db.UniversalModel
	Type        walletApi.TransactionType   `json:"type" gorm:"type:varchar(25);not null"`
	Status      walletApi.TransactionStatus `json:"status" gorm:"type:varchar(25);not null"`
	Amount      int64                       `json:"amount" gorm:"not null"`
	WalletID    uuid.UUID                   `json:"wallet_id" gorm:"type:uuid REFERENCES wallets(id)"`
	Description string                      `json:"description" gorm:"not null"`
	RefID       string                      `json:"ref_id" gorm:"not null; type:uuid"`
}
