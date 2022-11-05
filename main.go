package main

import (
	"fmt"
	"log"
	"net/http"

	app "example/rest/App"
	Domain "example/rest/Domain"
	"example/rest/Services"

	"github.com/gorilla/mux"
)

const webPort = ":8080"

func main() {
	fmt.Println("Starting App")
	app.Start()

	var router = mux.NewRouter()

	customerRepo := Domain.NewCustomerRepositoryDB()
	customerServices := Services.NewCustomerService(customerRepo)

	var customerHandlers = app.Customerhandlers{customerServices}

	router.HandleFunc("/customers", customerHandlers.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", customerHandlers.FindById).Methods("GET")

	fmt.Println("Starting Web Server on port", webPort)
	log.Fatal(http.ListenAndServe(webPort, router))

}
