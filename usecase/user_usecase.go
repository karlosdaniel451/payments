package usecase

import (
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	"github.com/karlosdaniel451/go-rest-api-template/repository"
)

type UserUseCase interface {
	Create(user *model.User) (*model.User, error)
	GetById(id uint) (*model.User, error)
	GetByFullName(fullName string) ([]*model.User, error)
	GetByEmailAddress(emailAddress string) (*model.User, error)
	GetByCPF(CPF string) (*model.User, error)
	GetByCNPJ(CNPJ string) (*model.User, error)
	DeleteById(id uint) error
	GetAll() ([]*model.User, error)
}

type UserUseCaseImpl struct {
	repository repository.UserRepository
}

func NewUserUseCaseImpl(
	repository repository.UserRepository,
) UserUseCaseImpl {

	return UserUseCaseImpl{repository: repository}
}

func (useCase UserUseCaseImpl) Create(user *model.User) (*model.User, error) {
	return useCase.repository.Create(user)
}

func (useCase UserUseCaseImpl) GetById(id uint) (*model.User, error) {
	return useCase.repository.GetById(id)
}

func (useCase UserUseCaseImpl) GetByFullName(fullName string) ([]*model.User, error) {
	return useCase.repository.GetByFullName(fullName)
}

func (useCase UserUseCaseImpl) GetByEmailAddress(emailAddress string) (*model.User, error) {
	return useCase.repository.GetByEmailAddress(emailAddress)
}

func (useCase UserUseCaseImpl) GetByCPF(CPF string) (*model.User, error) {
	return useCase.repository.GetByCPF(CPF)
}

func (useCase UserUseCaseImpl) GetByCNPJ(CNPJ string) (*model.User, error) {
	return useCase.repository.GetByCNPJ(CNPJ)
}

func (useCase UserUseCaseImpl) DeleteById(id uint) error {
	return useCase.repository.DeleteById(id)
}

func (useCase UserUseCaseImpl) GetAll() ([]*model.User, error) {
	return useCase.repository.GetAll()
}
