package usecase

import (
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	"github.com/karlosdaniel451/go-rest-api-template/repository"
)

type TransactionUseCase interface {
	Create(transaction *model.Transaction) (*model.Transaction, error)
	GetById(id uint) (*model.Transaction, error)
	DeleteById(id uint) error
	GetByPayerId(payerId uint) ([]*model.Transaction, error)
	GetByPayeeId(payeeId uint) ([]*model.Transaction, error)
	GetAll() ([]*model.Transaction, error)
}

type TransactionUseCaseImpl struct {
	repository repository.TransactionRepository
}

func NewTransactionUseCaseImpl(
	repository repository.TransactionRepository,
) TransactionUseCaseImpl {

	return TransactionUseCaseImpl{repository: repository}
}

func (useCase TransactionUseCaseImpl) Create(
	transaction *model.Transaction,
) (*model.Transaction, error) {

	return useCase.repository.Create(transaction)
}

func (useCase TransactionUseCaseImpl) GetById(id uint) (*model.Transaction, error) {
	return useCase.repository.GetById(id)
}

func (useCase TransactionUseCaseImpl) DeleteById(id uint) error {
	return useCase.repository.DeleteById(id)
}

func (useCase TransactionUseCaseImpl) GetByPayerId(id uint) ([]*model.Transaction, error) {
	return useCase.repository.GetByPayerId(id)
}

func (useCase TransactionUseCaseImpl) GetByPayeeId(id uint) ([]*model.Transaction, error) {
	return useCase.repository.GetByPayeeId(id)
}

func (useCase TransactionUseCaseImpl) GetAll() ([]*model.Transaction, error) {
	return useCase.repository.GetAll()
}
