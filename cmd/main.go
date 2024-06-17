package main

import (
	"database/sql"
	"doctorsFinal/internal"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	users   *internal.UserModel
	doctors *internal.DoctorModel
}

func main() {

	connStr := "user=postgres dbname=doctorsFinalFX password=cityzens sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		slog.Error("Cannot connect to the database")
	}
	slog.Info("Successfully connected to the database")
	defer db.Close()

	if err := db.Ping(); err != nil {
		slog.Error("Cannot ping the database")
	}

	app := &application{
		users:   &internal.UserModel{DB: db},
		doctors: &internal.DoctorModel{DB: db},
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("Let him cook on :" + port)

	if err := http.ListenAndServe(":"+port, app.routes()); err != nil {
		slog.Error("Cannot cook the server")
	}
}


func