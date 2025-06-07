package service

import (
	"github.com/theisaachome/go-eWallet/domain"
	"github.com/theisaachome/go-eWallet/dto"
	"github.com/theisaachome/go-eWallet/exception"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, error)
	GetCustomerByID(string) (*dto.CustomerResponse, *exception.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, error) {
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultCustomerService) GetCustomerByID(id string) (*dto.CustomerResponse, *exception.AppError) {

	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
