package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/srisudarshanrg/go-react-formula-one-backend/internal/models"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	type HomeData struct {
		CurrentDrivers []*models.CurrentDrivers `json:"current_drivers"`
		CurrentTeams   []*models.CurrentTeams   `json:"current_teams"`
		CurrentTracks  []*models.CurrentTracks  `json:"current_tracks"`
		Ok             string                   `json:"ok"`
	}

	drivers, err := app.GetCurrentDrivers()
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("coudn't get drivers from database"), http.StatusBadRequest)
		return
	}

	teams, err := app.GetCurrentTeams()
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("couldn't get teams from database"), http.StatusBadRequest)
		return
	}

	tracks, err := app.GetCurrentTracks()
	if err != nil {
		log.Println(err)
		app.errorJSON(w, errors.New("couldn't get tracks from database"), http.StatusBadRequest)
	}

	data := HomeData{
		CurrentDrivers: drivers,
		CurrentTeams:   teams,
		CurrentTracks:  tracks,
		Ok:             "Data is being sent from api successfully",
	}
	app.writeJSON(w, http.StatusOK, data)
}

func (app *Application) Drivers(w http.ResponseWriter, r *http.Request) {
	drivers, _ := app.GetCurrentDrivers()
	data := struct {
		Data interface{} `json:"data"`
	}{
		Data: drivers,
	}
	app.writeJSON(w, http.StatusOK, data)
}
