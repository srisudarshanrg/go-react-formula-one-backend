package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/srisudarshanrg/go-react-formula-one-backend/cmd/api"
)

const port = 10000

var app api.Application

func main() {
	app.DevelopmentFrontendLink = "http://localhost:3001"
	app.ProductionFrontendLink = "https://raptorf1.sudarshanraptor.world"
	app.DatabaseDSN = fmt.Sprintf("host=postgresql-raptor.alwaysdata.net port=5432 dbname=raptor_formula_one user=raptor password=%s", api.DatabasePassword)

	conn, err := app.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.Database = conn
	log.Println("Connected to Postgres Database")

	log.Println("Application starting on localhost:", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.Routes())
	if err != nil {
		log.Fatal(err)
	}
}
