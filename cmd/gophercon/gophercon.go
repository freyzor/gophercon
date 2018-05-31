package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/freyzor/gophercon/pkg/routing"
	"github.com/freyzor/gophercon/pkg/webserver"
	"github.com/freyzor/gophercon/version"
)

func main() {
	log.Printf("Service is staring.. %s %s %s", version.BuildTime, version.Commit, version.Release)

	shutdown := make(chan error, 2)

	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port was not set.")
	}

	r := routing.BaseRouter()
	ws := webserver.New("127.0.0.1", port, r)
	go func() { log.Fatal(ws.Start()) }()

	internalPort := os.Getenv("INTERNAL_PORT")
	if len(internalPort) == 0 {
		log.Fatal("Internal port was not set.")
	}

	diagnosticsRouter := routing.DiagnosticsRouter()
	diagnosticsServer := webserver.New("127.0.0.1", internalPort, diagnosticsRouter)
	go func() { log.Fatal(diagnosticsServer.Start()) }()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case killSignal := <-interrupt:
		log.Printf("Got %s. Stopping...", killSignal)
	case err := <-shutdown:
		log.Printf("Got an error '%s'. Stopping...", err)
	}

	err := ws.Stop()
	if err != nil {
		log.Print(err)
	}

	err = diagnosticsServer.Stop()
	if err != nil {
		log.Print(err)
	}

	// stop extra tasks ...
	log.Print("Service was stoped.")
}
