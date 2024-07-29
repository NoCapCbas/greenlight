package main

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
  // init http router instance
  router := httprouter.New()
  
  // register methods
  router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
  router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
  router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

  return router

}
