package main

import (
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func (app *application) Drivers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Drivers")
}
