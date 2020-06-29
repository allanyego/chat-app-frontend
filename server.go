package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	defaultPort := "3000"

	http.Handle("/", http.FileServer(http.Dir(".")))
	if !(port == "") {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
	}
}
