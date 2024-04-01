package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	"simplebank/db"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	fatalIfError("cannot connect to the db: ", err)

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	fatalIfError("cannot start server", err)
}

func fatalIfError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
