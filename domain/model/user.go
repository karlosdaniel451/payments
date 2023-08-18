package model

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type UserType string

const (
	Common     UserType = "COMMON"
	ShopKeeper UserType = "SHOP_KEEPER"
)

func (userType *UserType) Scan(value any) error {
	*userType = UserType(value.([]byte))
	return nil
}

func (userType UserType) Value() (driver.Value, error) {
	return string(userType), nil
}

type User struct {
	gorm.Model
	UserType         UserType      `json:"user_type" sql:"type:ENUM('COMMON','SHOP_KEEPER')" gorm:"column:user_type"`
	FullName         string        `json:"full_name" gorm:"varchar(200); not null"`
	CPF              string        `json:"cpf" gorm:"char(11); column:cpf; uniqueIndex"`
	CNPJ             string        `json:"cnpj" gorm:"char(14); column:cnpj; uniqueIndex"`
	EmailAddress     string        `json:"email_address" gorm:"varchar(320); not null; uniqueIndex"`
	PaymentsSent     []Transaction `json:"payments_sent" gorm:"foreignKey:PayerId"`
	PaymentsReceived []Transaction `json:"payments_received" gorm:"foreignKey:PayeeId"`
}
