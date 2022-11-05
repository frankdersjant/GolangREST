package app

import (
	"encoding/json"
	Services "example/rest/Services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//We define the handlers by wrapping them in a struct
//Customerhandlers in struct and we create the receiver functions (handlers)
//Which contains the customer service
type Customerhandlers struct {
	service Services.CustomerService
}

func (ch *Customerhandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
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
