package main

import (
	"fmt"
	"log"
	"net/http"

	app "example/rest/App"
	Domain "example/rest/Domain"
	Services "example/rest/Services"

	"github.com/gorilla/mux"
)

const webPort = ":8080"

func main() {
	fmt.Println("Starting App")
	app.Start()

	var router = mux.NewRouter()

	Services.NewCustomerService(Domain.NewCustomerRepositoryDB())

	//ch := app.Customerhandlers{NewCustomerService(Domain.NewCustomerRepositoryDB())}

	//router.HandleFunc("/customers", ch.GetAllCustomers).Methods("GET")

	//router.HandleFunc("/customers/{id}", ch.FindById).Methods("GET")

	fmt.Println("Starting Web Server on port", webPort)
	log.Fatal(http.ListenAndServe(webPort, router))

}
