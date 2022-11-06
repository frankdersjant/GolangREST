package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	app "example/rest/App"
	"example/rest/Domain"
	"example/rest/Services"

	"github.com/gorilla/mux"
)

const webPort = ":8080"

func main() {
	fmt.Println("Starting App")
	//app.Start()

	fmt.Println("Connectiong to db")

	conn, err := sql.Open("mysql", "root:change-me@tcp(localhost:3306)/usersdb")
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	for conn.Ping() != nil {
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Inserting user")

	_, err = conn.Exec("INSERT INTO users (id, name, email) VALUES (12, 'John', 'john@john.nl')")

	if err != nil {
		panic(err.Error())
	}

	var router = mux.NewRouter()

	customerRepo := Domain.NewCustomerRepositoryDB()
	customerServices := Services.NewCustomerService(customerRepo)

	var customerHandlers = app.Customerhandlers{customerServices}

	router.HandleFunc("/customers", customerHandlers.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", customerHandlers.FindById).Methods("GET")

	fmt.Println("Starting Web Server on port", webPort)
	log.Fatal(http.ListenAndServe(webPort, router))

}
