package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// define own multiplexer
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// define route
	// GET
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// // adding :[0-9]+ to tell the mux that only numeric value is accepted
	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	// // POST
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	// starting server
	log.Fatal(http.ListenAndServe(":8000", router))
}
