package entity

import "github.com/h-varmazyar/wallet/internal/pkg/db"

type Wallet struct {
	db.UniversalModel
	PhoneNumber  string         `json:"phone_number unique" gorm:"not null"`
	Amount       int64          `json:"amount" gorm:"not null"`
	Transactions []*Transaction `json:"transactions" gorm:"->;foreignkey:WalletID;references:ID"`
}
