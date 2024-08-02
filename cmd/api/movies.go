package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/NoCapCbas/greenlight/internal/data"
  "github.com/NoCapCbas/greenlight/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

  var input struct {
    Title string `json:"title"`
    Year int32 `json:"year"`
    Runtime data.Runtime `json:"runtime"`
    Genres []string `json:"genres`
  }

  err := app.readJSON(w, r, &input)
  if err != nil {
    app.badRequestResponse(w, r, err)
    return
  }
  
  movie := &data.Movie{
    Title: input.Title,
    Year: input.Year,
    Runtime: input.Runtime,
    Genres: input.Genres,
  }

  // initialize validator
  v := validator.New()

  if data.ValidateMovie(v, movie); !v.Valid() {
    // return response if validation failed
    app.failedValidationResponse(w, r, v.Errors)
    return
  }

  fmt.Fprintf(w, "%+v\n", input)


}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
  
  id, err := app.readIDParam(r) 
  if err != nil {
    app.notFoundResponse(w, r)
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
    app.serverErrorResponse(w, r, err)
  }

}
