package domain

import (
	"github.com/theisaachome/go-eWallet/dto"
	"github.com/theisaachome/go-eWallet/exception"
)

type Customer struct {
	Id          string `db:"id"`
	CustomerId  string `db:"customer_id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Phone       string `db:"phone"`
	Address     string `db:"address"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		CustomerId:  c.CustomerId,
		Name:        c.Name,
		Email:       c.Email,
		Phone:       c.Phone,
		Address:     c.Address,
		DateOfBirth: c.DateOfBirth,
		Status:      c.Status,
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, error)
	FindById(string) (*Customer, *exception.AppError)
}
