package main

import (
	"log"
	"os"

	"github.com/freyzor/gophercon/pkg/routing"
	"github.com/freyzor/gophercon/pkg/webserver"
)

func main() {
	log.Printf("Service is staring..")

	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port was not set.")
	}
	r := routing.BaseRouter()
	ws := webserver.New("127.0.0.1", port, r)
	log.Fatal(ws.Start())
	log.Printf("Service is stopping..")
}
