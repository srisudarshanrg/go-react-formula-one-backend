package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/srisudarshanrg/go-react-formula-one-backend/internal/models"
)

func (app *Application) search(w http.ResponseWriter, r *http.Request) {
	type searchRequestPayload struct {
		SearchQuery string `json:"search_query"`
	}
	var payload searchRequestPayload

	err := app.readJSON(r, &payload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	log.Println(payload.SearchQuery)

	type SearchData struct {
		DriverSearch []models.Driver        `json:"driver_search"`
		TeamsSearch  []models.AllTeams      `json:"team_search"`
		TracksSearch []models.CurrentTracks `json:"track_search"`
		OK           string                 `json:"ok"`
	}

	drivers, teams, tracks, err := app.SearchDB(payload.SearchQuery)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	var searchData = SearchData{
		DriverSearch: drivers,
		TeamsSearch:  teams,
		TracksSearch: tracks,
		OK:           "Search data sent successfully from api!",
	}

	err = app.writeJSON(w, http.StatusOK, searchData)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

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
		Ok:             "Data is being sent from api successfully!",
	}
	err = app.writeJSON(w, http.StatusOK, data)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *Application) Drivers(w http.ResponseWriter, r *http.Request) {
	type DriversData struct {
		CurrentDrivers       []*models.CurrentDrivers
		DriversChampionships []*models.Driver
		DriversWins          []*models.Driver
		DriversPodiums       []*models.Driver
		DriversPoles         []*models.Driver
	}

	currentDrivers, err := app.GetCurrentDrivers()
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	driversChampionships, err := app.GetDriversByAchievement("championships")
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	driversWins, err := app.GetDriversByAchievement("wins")
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	driversPodiums, err := app.GetDriversByAchievement("podiums")
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	driversPoles, err := app.GetDriversByAchievement("pole_positions")
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	data := DriversData{
		CurrentDrivers:       currentDrivers,
		DriversChampionships: driversChampionships,
		DriversWins:          driversWins,
		DriversPodiums:       driversPodiums,
		DriversPoles:         driversPoles,
	}

	err = app.writeJSON(w, http.StatusOK, data)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}
