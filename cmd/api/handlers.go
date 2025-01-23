package api

import (
	"fmt"
	"net/http"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func (app *Application) Drivers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Drivers")
}
