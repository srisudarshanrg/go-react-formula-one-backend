package api

import "database/sql"

type Application struct {
	FrontendLink string
	DatabaseDSN  string
	Database     *sql.DB
}
