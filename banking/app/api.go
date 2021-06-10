package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neoyewchuan/RestDevGo/banking/service"
)

type customer struct {
	ID      string `json:"customer-id" xml:"customer-id"`
	Name    string `json:"full-name" xml:"full-name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip-code" xml:"zip-code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []customer{
	// 	{ID: "001", Name: "Ashish", City: "New Delhi", ZipCode: "110075"},
	// 	{ID: "002", Name: "Rob", City: "New Delhi", ZipCode: "110075"},
	// 	{ID: "003", Name: "David Parker", City: "New York City", ZipCode: "11375"},
	// 	{ID: "007", Name: "James Bond", City: "San Monique", ZipCode: "00007"}
	// }
	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	//fmt.Fprint(w, vars["customer_id"])
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, err.Error())
	} else {

		if r.Header.Get("Content-Type") == "application/xml" {
			w.Header().Add("Content-type", "application/xml")
			xml.NewEncoder(w).Encode(customer)
		} else {
			w.Header().Add("Content-type", "application/json")
			json.NewEncoder(w).Encode(customer)
		}
	}

}

// func (ch CustomerHandlers) createCustomer(w http.ResponseWriter, r *http.Request) {
// 	var c customer

// 	_ = json.NewDecoder(r.Body).Decode(&c)

// 	customers = append(customers, c)

// 	json.NewEncoder(w).Encode(&c)

// 	//w.WriteHeader(http.StatusOK)
// }
