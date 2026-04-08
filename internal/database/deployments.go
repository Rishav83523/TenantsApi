package database

import (
	"database/sql"
)

type DeploymentModel struct {
	DB *sql.DB
}


type deployment struct {
	ID int
ServiceID int
Version string
Status string
Environment string
DeployedAt string	
}