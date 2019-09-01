package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"stash-mono-repo/service/usersvc"
	"syscall"

	_ "github.com/lib/pq"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	// our usersvc service
	srv := usersvc.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := usersvc.Endpoints{
		StatusEndpoint:   usersvc.MakeStatusEndpoint(*srv),
		ApprovalEndpoint: usersvc.MakeApprovalEndpoint(*srv),
	}

	// HTTP transport
	go func() {
		log.Println("approvalsvc is listening on port:", *httpAddr)
		handler := usersvc.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
