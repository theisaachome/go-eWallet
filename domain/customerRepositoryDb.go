package domain

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/theisaachome/go-eWallet/exception"
	"github.com/theisaachome/go-eWallet/logger"
	"os"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, error) {
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAlQuery := "SELECT id,customer_id,name,email,phone,address,date_of_birth,status FROM customers"
		err = d.client.Select(&customers, findAlQuery)

	} else {
		findAlQuery := "SELECT id,customer_id,name,email,phone,address,date_of_birth,status FROM customers WHERE status=?"
		err = d.client.Select(&customers, findAlQuery, status)
	}
	if err != nil {
		logger.Error("Error while fetching customers table rows:" + err.Error())
		return nil, err
	}
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *exception.AppError) {
	customerQuery := "SELECT id,customer_id,name,email,phone,address,date_of_birth,status FROM customers WHERE id=?"
	var c Customer
	err := d.client.Get(&c, customerQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, exception.NewNotFoundException("customer not found")
		} else {
			logger.Error("Error while Scanning customer by id in customer table rows:" + err.Error())
			return nil, exception.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbUsr := os.Getenv("DB_USR")
	dbPassWd := os.Getenv("DB_PASSWD")
	dbHost := os.Getenv("DB_HOST")
	dbPORT := os.Getenv("DB_PORT")
	dbNAME := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsr, dbPassWd, dbHost, dbPORT, dbNAME)
	// ...s
	db, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDb{db}
}
