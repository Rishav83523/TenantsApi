package main

import (
	"database/sql"
	"log"
	"rest-api-in-gin/internal/database"

	_ "modernc.org/sqlite"
)


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