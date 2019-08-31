package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"stash-mono-repo/service/approvalsvc"
	"syscall"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "apple"
	dbname   = "Approval"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	/*sqlStatement := `
	INSERT INTO approval (id, state, comment, priority, service, permission, deadline)
	VALUES ('123',2,'abc',0,'123',3)`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		//		panic(err)
	}*/

	var (
		httpAddr = flag.String("http", ":8000", "http listen address")
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
		GetApprovalsEndpoint:   approvalsvc.MakeGetApprovalsEndpoint(srv),
		AddApprovalEndpoint:    approvalsvc.MakeAddApprovalEndpoint(srv),
		UpdateApprovalEndpoint: approvalsvc.MakeUpdateApprovalEndpoint(srv),
		StatusEndpoint:         approvalsvc.MakeStatusEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("approvalsvc is listening on port:", *httpAddr)
		handler := approvalsvc.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
