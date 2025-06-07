package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/theisaachome/go-eWallet/domain"
	"github.com/theisaachome/go-eWallet/service"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable SERVER_ADDRESS and SERVER_PORT not set")
	}
}

func StartApplication() {
	sanityCheck()
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	// wiring
	//handlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	handlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", handlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", handlers.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
