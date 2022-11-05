package app

import (
	"encoding/json"
	CustomerServices "example/rest/Services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Customerhandlers is a struct
//Which contains the customer service
type Customerhandlers struct {
	service CustomerServices.CustomerService
}

func (ch *Customerhandlers) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customer, err := ch.service.FindCustomerById(params["id"])
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func (ch *Customerhandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
