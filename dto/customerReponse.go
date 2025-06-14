package dto

type CustomerResponse struct {
	Id          string `json:"id"`
	CustomerId  string `json:"customer_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}
