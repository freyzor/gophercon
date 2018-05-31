package main

import (
	"log"
	"net/http"
	"os"

	"github.com/freyzor/gophercon/pkg/routing"
)

func main() {
	log.Printf("Service is staring..")

	port := os.Getenv("GOPHERCON_SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port was not set.")
	}
	r := routing.BaseRouter()
	log.Fatal(http.ListenAndServe(":"+port, r))
	log.Printf("Service is stopping..")
}
