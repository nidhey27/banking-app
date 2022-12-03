package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/ndihey27/banking-app/api"
	db "github.com/ndihey27/banking-app/db/sql"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:mypassword@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
}
