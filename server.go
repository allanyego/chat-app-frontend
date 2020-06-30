package main

// Module to serve the production build folder

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Get PORT set by Heroku
	port := os.Getenv("PORT")
	defaultPort := "3000"

	http.Handle("/", http.FileServer(http.Dir(".")))
	if !(port == "") {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
	}
}
