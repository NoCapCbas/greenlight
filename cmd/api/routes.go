package main

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
  // init http router instance
  router := httprouter.New()
  // Covert notFoundResponse() helper to a http.Handler using http.HandlerFunc() adapter, and then set it as a custom error handler for 404 responses
  router.NotFound = http.HandlerFunc(app.notFoundResponse)
  // Covert MethodNotAllowedResponse() helper to a http.Handler using http.HandlerFunc() adapter, and then set it as a custom error handler for 405 responses
  router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
  
  // register methods
  router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
  router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
  router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
  
  // wrap router with panic recovery middleware
  return app.recoverPanic(router)

}
