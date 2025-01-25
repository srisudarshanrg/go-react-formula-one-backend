package models

// Driver is the model for a driver object in the database
type Driver struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Age               int         `json:"age"`
	Nationality       string      `json:"nationality"`
	PolePositions     int         `json:"pole_positions"`
	Podiums           int         `json:"podiums"`
	Wins              int         `json:"wins"`
	Championships     int         `json:"championships"`
	YearsDriven       string      `json:"years_driven"`
	TeamsDriven       string      `json:"teams_driven"`
	NumberYearsDriven int         `json:"number_years_driven"`
	CreatedAt         interface{} `json:"-"`
	UpdatedAt         interface{} `json:"-"`
}

// CurrentDrivers is the model for a current drivers object in the database
type CurrentDrivers struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	Number             int         `json:"number"`
	Position           int         `json:"position"`
	Points             int         `json:"points"`
	Team               string      `json:"team"`
	TeamColor          string      `json:"team_color"`
	PercentagePoints   float64     `json:"percentage_points"`
	ChampionshipWinner bool        `json:"championship_winner"`
	CreatedAt          interface{} `json:"-"`
	UpdatedAt          interface{} `json:"-"`
}

// CurrentTeams is the model for a current teams object in the database
type CurrentTeams struct {
	ID                   int         `json:"id"`
	Name                 string      `json:"name"`
	Drivers              string      `json:"drivers"`
	TotalPoints          int         `json:"total_points"`
	ConstructorsPosition int         `json:"constructors_position"`
	HighestPointsHaul    int         `json:"highest_points_haul"`
	ChampionshipWinner   bool        `json:"championship_winner"`
	LogoLink             string      `json:"logo_link"`
	CreatedAt            interface{} `json:"-"`
	UpdatedAt            interface{} `json:"-"`
}

// AllTeams is the model for an all teams object in the database
type AllTeams struct {
	ID                         int         `json:"id"`
	Name                       string      `json:"name"`
	Nationality                string      `json:"nationality"`
	YearJoined                 int         `json:"year_joined"`
	Poles                      int         `json:"poles"`
	Podiums                    int         `json:"podiums"`
	Wins                       int         `json:"wins"`
	Championships              int         `json:"championships"`
	NotableCars                string      `json:"notable_cars"`
	ChampionshipWinningDrivers string      `json:"championship_winning_drivers"`
	CreatedAt                  interface{} `json:"-"`
	UpdatedAt                  interface{} `json:"-"`
}

// CurrentTracks is the model for a current track object in the database
type CurrentTracks struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	Length          int         `json:"length"`
	NumberCorners   int         `json:"number_corners"`
	NumberStraights int         `json:"number_straights"`
	NumberDRSZones  int         `json:"number_drs_zones"`
	Year            int         `json:"year"`
	Country         string      `json:"country"`
	Image           string      `json:"image"`
	CreatedAt       interface{} `json:"-"`
	UpdatedAt       interface{} `json:"-"`
}
