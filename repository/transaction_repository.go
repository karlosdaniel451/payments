package repository

import (
	"fmt"

	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	"github.com/karlosdaniel451/go-rest-api-template/errs"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) (*model.Transaction, error)
	GetById(id uint) (*model.Transaction, error)
	DeleteById(id uint) error
	GetByPayerId(payerId uint) ([]*model.Transaction, error)
	GetByPayeeId(payeeId uint) ([]*model.Transaction, error)
	GetAll() ([]*model.Transaction, error)
}

type TransactionRepositoryDB struct {
	db *gorm.DB
}

func NewTransactionRepositoryDB(db *gorm.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{db: db}
}

func (repository TransactionRepositoryDB) Create(
	transaction *model.Transaction,
) (*model.Transaction, error) {

	result := repository.db.Create(transaction)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("error when inserting transaction: %s", result.Error)
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return transaction, nil
}

func (repository TransactionRepositoryDB) GetById(id uint) (*model.Transaction, error) {
	var transaction model.Transaction

	result := repository.db.First(&transaction, "id = ?", id)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errs.NotFoundError{
				Message: fmt.Sprintf("there is no transaction with id %d", id),
			}
		}
		return nil, result.Error
	}

	return &transaction, nil
}

func (repository TransactionRepositoryDB) DeleteById(id uint) error {
	var user model.Transaction

	result := repository.db.First(&user, id)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return errs.NotFoundError{
				Message: fmt.Sprintf("there is no transaction with id %d", id),
			}
		}
		return result.Error
	}
	result = result.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository TransactionRepositoryDB) GetByPayerId(payerId uint) ([]*model.Transaction, error) {
	allUsers := make([]*model.Transaction, 0)

	result := repository.db.Find(&allUsers, "payerId = ?", payerId)
	if result.Error != nil {
		return nil, result.Error
	}

	return allUsers, nil
}

func (repository TransactionRepositoryDB) GetByPayeeId(payeeId uint) ([]*model.Transaction, error) {
	allUsers := make([]*model.Transaction, 0)

	result := repository.db.Find(&allUsers, "payee_id = ?", payeeId)
	if result.Error != nil {
		return nil, result.Error
	}

	return allUsers, nil
}

func (repository TransactionRepositoryDB) GetAll() ([]*model.Transaction, error) {
	allUsers := make([]*model.Transaction, 0)

	result := repository.db.Find(&allUsers)
	if result.Error != nil {
		return nil, result.Error
	}

	return allUsers, nil
}
