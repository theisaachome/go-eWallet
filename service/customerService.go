package service

import (
	"github.com/theisaachome/go-eWallet/domain"
	"github.com/theisaachome/go-eWallet/exception"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomerByID(string) (*domain.Customer, *exception.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(id string) (*domain.Customer, *exception.AppError) {
	return s.repo.FindById(id)
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
