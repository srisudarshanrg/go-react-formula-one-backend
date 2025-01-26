package api

import "github.com/srisudarshanrg/go-react-formula-one-backend/internal/models"

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
	var drivers []*models.Driver

	return drivers, nil
}

// GetTeamsByAchievement gets all teams by decreasing order of number of a specific achievement passed as a parameter
func (app *Application) GetTeamsByAchievement(achievementName string) ([]*models.AllTeams, error) {
	var teams []*models.AllTeams

	return teams, nil
}
