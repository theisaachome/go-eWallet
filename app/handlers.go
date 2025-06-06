package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/theisaachome/go-eWallet/service"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{"Isaac", "Ho chi Minh", "11000"},
	//	{"John Peter", "Da Nang", "13000"},
	//	{"Olivia", "Hanoi", "12000"},
	//	{"Paul", "sapa", "14000"},
	//	{"Mary", "Cho tha", "17000"},
	//}
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		// xml format
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		// json format response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
