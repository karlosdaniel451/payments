package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Value   decimal.Decimal `json:"value" gorm:"decimal(20, 2); not null"`
	PayerId uint            `json:"payer_id" gorm:"index:idx_transactions; not null"`
	PayeeId uint            `json:"payee_id" gorm:"index:idx_transactions; not null"`
}
