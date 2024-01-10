package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go movies up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println()
	}
	w.Header().Set("Connect-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "2023-03-7")

	spiderman := models.Movie{
		ID:           1,
		Title:        "Spiderman",
		ReleaseDate:  rd,
		MPAARating:   "R",
		RunTime:      120,
		Description:  "Something excellent",
		CreatedAt:    time.Now(),
		UpdatedField: time.Now(),
	}

	batman := models.Movie{
		ID:           2,
		Title:        "Batman",
		ReleaseDate:  rd,
		MPAARating:   "PG-13",
		RunTime:      120,
		Description:  "Very cool movie excellent",
		CreatedAt:    time.Now(),
		UpdatedField: time.Now(),
	}

	movies = append(movies, spiderman)
	movies = append(movies, batman)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println()
	}
	w.Header().Set("Connect-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
