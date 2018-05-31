package main

import (
	"log"
	"os"

	"github.com/freyzor/gophercon/pkg/routing"
	"github.com/freyzor/gophercon/pkg/webserver"
	"github.com/freyzor/gophercon/version"
)

func main() {
	log.Printf("Service is staring.. %s %s %s", version.BuildTime, version.Commit, version.Release)

	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port was not set.")
	}
	r := routing.BaseRouter()
	ws := webserver.New("127.0.0.1", port, r)
	log.Fatal(ws.Start())
	log.Printf("Service is stopping..")
}
