package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/neoyewchuan/RestDevGo/banking/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	var status string
	keys, ok := r.URL.Query()["status"]
	if !ok || len(keys[0]) < 1 {
		status = ""
	} else {
		status = keys[0]
	}
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		encodeResponse(w, r, err.Code, err.AsMessage())
	} else {
		encodeResponse(w, r, http.StatusOK, customers)
	}

	// w.Header().Add("Content-type", r.Header.Get("Content-Type"))
	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	xml.NewEncoder(w).Encode(customers)
	// } else {
	// 	json.NewEncoder(w).Encode(customers)
	// }
}

func (ch CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	//fmt.Fprint(w, vars["customer_id"])
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		encodeResponse(w, r, err.Code, err.AsMessage())
		// w.Header().Add("Content-type", r.Header.Get("Content-Type"))
		// w.WriteHeader(err.Code)
		// if r.Header.Get("Content-Type") == "application/xml" {
		// 	xml.NewEncoder(w).Encode(err.AsMessage())
		// } else {
		// 	json.NewEncoder(w).Encode(err.AsMessage())
		// }
	} else {
		// w.Header().Add("Content-type", r.Header.Get("Content-Type"))
		// w.WriteHeader(http.StatusOK)
		// if r.Header.Get("Content-Type") == "application/xml" {
		// 	xml.NewEncoder(w).Encode(customer)
		// } else {
		// 	json.NewEncoder(w).Encode(customer)
		// }
		encodeResponse(w, r, http.StatusOK, customer)
	}
}

func encodeResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	contentType := "application/json"
	if r.Header.Get("Content-Type") != contentType && r.Header.Get("Content-Type") != "" {
		contentType = r.Header.Get("Content-Type")
	}
	w.Header().Add("Content-Type", contentType)
	w.WriteHeader(code)
	if contentType == "application/xml" {
		if err := xml.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	}
}

// func (ch CustomerHandler) createCustomer(w http.ResponseWriter, r *http.Request) {
// 	var c customer

// 	_ = json.NewDecoder(r.Body).Decode(&c)

// 	customers = append(customers, c)

// 	json.NewEncoder(w).Encode(&c)

// 	//w.WriteHeader(http.StatusOK)
// }
