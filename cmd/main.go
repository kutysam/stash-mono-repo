package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"approvalsvc"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	// our approvalsvc service
	srv := approvalsvc.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := approvalsvc.Endpoints{
		GetEndpoint:      approvalsvc.MakeGetEndpoint(srv),
		StatusEndpoint:   approvalsvc.MakeStatusEndpoint(srv),
		ValidateEndpoint: approvalsvc.MakeValidateEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("approvalsvc is listening on port:", *httpAddr)
		handler := approvalsvc.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
