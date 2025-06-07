package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAlQuery := "SELECT id,customer_id,name,email,phone,address,date_of_birth,status FROM customers"

	rows, err := d.client.Query(findAlQuery)
	if err != nil {
		log.Println("Error while fetching customers table rows:", err)
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.CustomerId, &c.Name, &c.Email, &c.Phone, &c.Address, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while fetching customers table rows:", err)
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, error) {
	customerQuery := "SELECT id,customer_id,name,email,phone,address,date_of_birth,status FROM customers WHERE id=?"
	row := d.client.QueryRow(customerQuery, id)
	var c Customer
	err := row.Scan(&c.Id, &c.CustomerId, &c.Name, &c.Email, &c.Phone, &c.Address, &c.DateOfBirth, &c.Status)
	if err != nil {
		log.Println("Error while Scanning customer by id in customer table rows:", err)
		return nil, err
	}
	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	// ...
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/db")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDb{db}
}
