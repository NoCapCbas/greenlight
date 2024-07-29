package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/NoCapCbas/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

  fmt.Fprintln(w, "create new movie")

}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
  
  id, err := app.readIDParam(r) 
  if err != nil {
    http.NotFound(w, r)
    return
  }
  // Create a new instance of the Movie struct, containing the ID we extracted from the URL and some dummy data  
  movie := data.Movie{
    ID: id,
    CreatedAt: time.Now(),
    Title: "Casablanca",
    Runtime: 102,
    Genres: []string{"drama", "romance", "war"},
    Version: 1,
  }

  // Encode the struct to json and send it as the http response.
  // wrap movie data in envelope
  err = app.writeJSON(w, http.StatusOK, envelope{"movie" : movie}, nil)
  if err != nil {
    app.logger.Error(err.Error())
    http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
  }

}
