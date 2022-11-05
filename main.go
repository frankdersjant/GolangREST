package main

import (
	"fmt"
	"log"
	"net/http"

	app "example/rest/App"

	"github.com/gorilla/mux"
)

const webPort = ":8080"

func main() {
	fmt.Println("Starting App")
	app.Start()

	var router = mux.NewRouter()

	fmt.Println("Starting Web Server on port", webPort)
	log.Fatal(http.ListenAndServe(webPort, router))

}
