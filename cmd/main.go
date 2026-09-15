package main

import (
	"log"
	"net/http"

	"github.com/ryansissom/go-rest-api/internal/router"
)

func main() {
	if err := http.ListenAndServe(":8080", router.New()); err != nil {
		log.Fatal("Failed to start server", err)
	}
}
