package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const port = 10000

type application struct {
	FrontendLink string
	DatabaseDSN  string
	Database     *sql.DB
}

var app application

func main() {
	app.FrontendLink = "http://localhost:3000"
	app.DatabaseDSN = fmt.Sprintf("host=postgresql-raptor.alwaysdata.net port=5432 dbname=raptor_formula_one user=raptor password=%s", databasePassword)

	conn, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.Database = conn

	log.Println(databasePassword)
	log.Println("eee")

	log.Println("Connected to Postgres Database")

	log.Println("Application starting on localhost", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), routes())
	if err != nil {
		log.Fatal(err)
	}
}
