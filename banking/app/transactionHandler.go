package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/neoyewchuan/RestDevGo/banking/dto"
	"github.com/neoyewchuan/RestDevGo/banking/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (th TransactionHandler) newTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		encodeResponse(w, r, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerID = id
		account, appErr := th.service.NewTransaction(request)
		if appErr != nil {
			encodeResponse(w, r, appErr.Code, appErr.Message)
		} else {
			encodeResponse(w, r, http.StatusCreated, account)
		}
	}
}
