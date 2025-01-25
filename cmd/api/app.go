package api

import "database/sql"

type Application struct {
	DevelopmentFrontendLink string
	ProductionFrontendLink  string
	DatabaseDSN             string
	Database                *sql.DB
}
