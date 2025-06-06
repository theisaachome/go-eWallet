package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from Greeting endpoint")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Isaac", "Ho chi Minh", "11000"},
		{"John Peter", "Da Nang", "13000"},
		{"Olivia", "Hanoi", "12000"},
		{"Paul", "sapa", "14000"},
		{"Mary", "Cho tha", "17000"},
	}

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
