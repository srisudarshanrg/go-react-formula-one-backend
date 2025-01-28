package api

import (
	"fmt"

	"github.com/srisudarshanrg/go-react-formula-one-backend/internal/models"
)

func (app *Application) SearchDB(searchQuery string) ([]models.Driver, []models.AllTeams, []models.CurrentTracks, error) {
	var drivers []models.Driver
	var teams []models.AllTeams
	var tracks []models.CurrentTracks

	searchQuery = "%" + searchQuery + "%"
	queryDrivers := `select * from drivers where lower(name) like $1`
	rowsDrivers, err := app.Database.Query(queryDrivers, searchQuery)
	if err != nil {
		return nil, nil, nil, err
	}
	defer rowsDrivers.Close()

	queryTeams := `select * from all_teams where lower(name) like $1`
	rowsTeams, err := app.Database.Query(queryTeams, searchQuery)
	if err != nil {
		return nil, nil, nil, err
	}
	defer rowsTeams.Close()

	queryTracks := `select * from current_tracks where lower(name) like $1`
	rowsTracks, err := app.Database.Query(queryTracks, searchQuery)
	if err != nil {
		return nil, nil, nil, err
	}
	defer rowsTracks.Close()

	for rowsDrivers.Next() {
		var driver models.Driver
		err = rowsDrivers.Scan(
			&driver.ID,
			&driver.Name,
			&driver.Age,
			&driver.Nationality,
			&driver.PolePositions,
			&driver.Podiums,
			&driver.Wins,
			&driver.Championships,
			&driver.YearsDriven,
			&driver.TeamsDriven,
			&driver.NumberYearsDriven,
			&driver.CreatedAt,
			&driver.UpdatedAt,
		)
		if err != nil {
			return nil, nil, nil, err
		}
		drivers = append(drivers, driver)
	}

	for rowsTeams.Next() {
		var team models.AllTeams
		err = rowsTeams.Scan(
			&team.ID,
			&team.Name,
			&team.Nationality,
			&team.YearJoined,
			&team.Poles,
			&team.Podiums,
			&team.Wins,
			&team.Championships,
			&team.NotableCars,
			&team.ChampionshipWinningDrivers,
			&team.CreatedAt,
			&team.UpdatedAt,
		)
		if err != nil {
			return nil, nil, nil, err
		}
		teams = append(teams, team)
	}

	for rowsTracks.Next() {
		var track models.CurrentTracks
		err = rowsTracks.Scan(
			&track.ID,
			&track.Name,
			&track.Length,
			&track.NumberCorners,
			&track.NumberStraights,
			&track.NumberDRSZones,
			&track.Year,
			&track.Country,
			&track.Image,
			&track.CreatedAt,
			&track.UpdatedAt,
		)
		if err != nil {
			return nil, nil, nil, err
		}
		tracks = append(tracks, track)
	}

	return drivers, teams, tracks, nil
}

// GetCurrentDrivers gets all the drivers of F1 2024 from the database and returns them as a list
func (app *Application) GetCurrentDrivers() ([]*models.CurrentDrivers, error) {
	query := `select * from current_drivers`
	rows, err := app.Database.Query(query)
	if err != nil {
		return nil, err
	}

	var drivers []*models.CurrentDrivers

	for rows.Next() {
		var driver models.CurrentDrivers
		err = rows.Scan(
			&driver.ID,
			&driver.Name,
			&driver.Number,
			&driver.Position,
			&driver.Points,
			&driver.Team,
			&driver.TeamColor,
			&driver.PercentagePoints,
			&driver.ChampionshipWinner,
			&driver.CreatedAt,
			&driver.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, &driver)
	}

	return drivers, nil
}

// GetCurrentTeams gets all the F1 2024 teams from the database and returns them as a list
func (app *Application) GetCurrentTeams() ([]*models.CurrentTeams, error) {
	query := `select * from current_teams`
	rows, err := app.Database.Query(query)
	if err != nil {
		return nil, err
	}

	var teams []*models.CurrentTeams

	for rows.Next() {
		var team models.CurrentTeams
		err = rows.Scan(
			&team.ID,
			&team.Name,
			&team.Drivers,
			&team.TotalPoints,
			&team.ConstructorsPosition,
			&team.HighestPointsHaul,
			&team.ChampionshipWinner,
			&team.LogoLink,
			&team.CreatedAt,
			&team.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		teams = append(teams, &team)
	}

	return teams, nil
}

// GetCurrentTracks gets all the F1 2024 tracks from the database and returns them as a list
func (app *Application) GetCurrentTracks() ([]*models.CurrentTracks, error) {
	query := `select * from current_tracks`
	rows, err := app.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []*models.CurrentTracks

	for rows.Next() {
		var driver models.CurrentTracks
		err = rows.Scan(
			&driver.ID,
			&driver.Name,
			&driver.Length,
			&driver.NumberCorners,
			&driver.NumberStraights,
			&driver.NumberDRSZones,
			&driver.Year,
			&driver.Country,
			&driver.Image,
			&driver.CreatedAt,
			&driver.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, &driver)
	}

	return tracks, nil
}

// GetDriversByAchievement gets all drivers by decreasing order of number of a specific achievement passed as a parameter
func (app *Application) GetDriversByAchievement(achievementName string) ([]*models.Driver, error) {
	query := fmt.Sprintf("select * from drivers where %s > 0 order by %s desc", achievementName, achievementName)
	rows, err := app.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []*models.Driver

	for rows.Next() {
		var driver models.Driver
		err = rows.Scan(
			&driver.ID,
			&driver.Name,
			&driver.Age,
			&driver.Nationality,
			&driver.PolePositions,
			&driver.Podiums,
			&driver.Wins,
			&driver.Championships,
			&driver.YearsDriven,
			&driver.TeamsDriven,
			&driver.NumberYearsDriven,
			&driver.CreatedAt,
			&driver.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		drivers = append(drivers, &driver)
	}

	return drivers, nil
}

// GetTeamsByAchievement gets all teams by decreasing order of number of a specific achievement passed as a parameter
func (app *Application) GetTeamsByAchievement(achievementName string) ([]*models.AllTeams, error) {
	var teams []*models.AllTeams

	query := fmt.Sprintf("select * from all_teams where %s > 0 order by %s desc", achievementName, achievementName)
	rows, err := app.Database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var team models.AllTeams
		err = rows.Scan(
			&team.ID,
			&team.Name,
			&team.Nationality,
			&team.YearJoined,
			&team.Poles,
			&team.Podiums,
			&team.Wins,
			&team.Championships,
			&team.NotableCars,
			&team.ChampionshipWinningDrivers,
			&team.CreatedAt,
			&team.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
	}

	return teams, nil
}
