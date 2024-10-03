package main

import (
	"log"
	"net/http"

	"github.com/MeherKandukuri/studioClasses_API/routes"
)

const portNumber = ":8080"

func main() {
	log.Println("We are starting on port number:", portNumber)
	server := run()
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

// setting up server with handler
func run() *http.Server {
	router := routes.Routes()
	return &http.Server{
		Addr:    portNumber,
		Handler: router,
	}
}
