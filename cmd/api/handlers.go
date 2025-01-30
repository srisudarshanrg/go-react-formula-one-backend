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

func (app *Application) compare(w http.ResponseWriter, r *http.Request) {

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
		CurrentDrivers       []*models.CurrentDrivers `json:"current_drivers"`
		DriversChampionships []models.Driver          `json:"drivers_championships"`
		DriversWins          []models.Driver          `json:"drivers_wins"`
		DriversPodiums       []models.Driver          `json:"drivers_podiums"`
		DriversPoles         []models.Driver          `json:"drivers_poles"`
		OK                   string                   `json:"ok"`
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
		OK:                   "Get data successfully sent from drivers backend api!",
	}

	err = app.writeJSON(w, http.StatusOK, data)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *Application) Teams(w http.ResponseWriter, r *http.Request) {
	type TeamsData struct {
		TeamsChampionships []models.AllTeams      `json:"teams_championships"`
		TeamsWins          []models.AllTeams      `json:"teams_wins"`
		TeamsPodiums       []models.AllTeams      `json:"teams_podiums"`
		TeamsPoles         []models.AllTeams      `json:"teams_poles"`
		CurrentTeams       []*models.CurrentTeams `json:"current_teams"`
		OK                 string                 `json:"ok"`
	}

	teamsChampionships, err := app.GetTeamsByAchievement("championships")
	if err != nil {
		log.Println(err)
		app.writeJSON(w, http.StatusBadRequest, err)
		return
	}

	teamsWins, err := app.GetTeamsByAchievement("wins")
	if err != nil {
		log.Println(err)
		app.writeJSON(w, http.StatusBadRequest, err)
		return
	}

	teamsPodiums, err := app.GetTeamsByAchievement("podiums")
	if err != nil {
		log.Println(err)
		app.writeJSON(w, http.StatusBadRequest, err)
		return
	}

	teamsPoles, err := app.GetTeamsByAchievement("poles")
	if err != nil {
		log.Println(err)
		app.writeJSON(w, http.StatusBadRequest, err)
		return
	}

	currentTeams, err := app.GetCurrentTeams()
	if err != nil {
		log.Println(err)
		app.writeJSON(w, http.StatusBadRequest, err)
		return
	}

	var data = TeamsData{
		TeamsChampionships: teamsChampionships,
		TeamsWins:          teamsWins,
		TeamsPodiums:       teamsPodiums,
		TeamsPoles:         teamsPoles,
		CurrentTeams:       currentTeams,
		OK:                 "Get data successfully sent from teams backend api!",
	}

	err = app.writeJSON(w, http.StatusOK, data)
	if err != nil {
		log.Println(err)
		app.writeJSON(w, http.StatusBadRequest, err)
		return
	}
}
