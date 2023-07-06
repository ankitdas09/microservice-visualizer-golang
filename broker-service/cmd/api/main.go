package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const PORT = "8080"

func main() {
	app := Config{}

	fmt.Printf("Broker service on PORT %s\n", PORT)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
