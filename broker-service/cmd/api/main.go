package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8080"

type Config struct{}

func main() {
	fmt.Printf("starting beoker service\n")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: routes(),
	}

	log.Fatal(server.ListenAndServe())
}
