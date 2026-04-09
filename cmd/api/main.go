package main

import (
	"database/sql"
	"log"
	"rest-api-in-gin/internal/database"

	_ "rest-api-in-gin/docs"

	_ "modernc.org/sqlite"
)

// @title Tenants API
// @version 1.0
// @description REST API for tenants, projects, services, and deployments.
// @BasePath /api/v1
// @schemes http
// @tag.name Tenants
// @tag.description Tenant management endpoints.
// @tag.name Projects
// @tag.description Project management endpoints.
// @tag.name Services
// @tag.description Service management endpoints.
// @tag.name Deployments
// @tag.description Deployment management endpoints.

type application struct { 
	port int
	models database.Models
}

func main() {
	db, err := sql.Open("sqlite","./data.db")
	if err != nil {
		log.Fatal(err);
	}
	defer db.Close()

	models := database.NewModels(db)
	app := &application{
	port: 4000,
	models: models,
	}

	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}