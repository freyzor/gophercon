package main

import (
	"log"
	"net/http"

	"github.com/freyzor/gophercon/pkg/routing"
)

func main() {
	log.Printf("Service is staring..")
	r := routing.BaseRouter()
	http.ListenAndServe(":8000", r)
	log.Printf("Service is stopping..")
}
