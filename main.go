package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/ndihey27/banking-app/api"
	db "github.com/ndihey27/banking-app/db/sql"
	"github.com/ndihey27/banking-app/utils"
)

func main() {

	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot Load Configuration: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
}
