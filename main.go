package main

import (
	"fmt"
	"log"
	"net/http"

	app "example/rest/App"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting App")
	app.Start()

	var router = mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
