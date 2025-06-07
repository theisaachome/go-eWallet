package app

import (
	"github.com/gorilla/mux"
	"github.com/theisaachome/go-eWallet/domain"
	"github.com/theisaachome/go-eWallet/service"
	"log"
	"net/http"
)

func StartApplication() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	// wiring
	//handlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	handlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", handlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", handlers.getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
