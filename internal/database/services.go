package database

import (
	"database/sql"
)

type ServiceModel struct { 
	DB *sql.DB
}

type service struct { 
	ID int
Name string
Type string
Language string
ProjectID int
CreatedAt string
UpdatedAt string
}