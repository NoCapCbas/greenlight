package main

import (
  "fmt"
  "net/http"
)

func (app *application) logError(r *http.Request, err error) {
  // generic helper for logging an error message along with the current request and url
  var (
    method = r.Method
    url = r.URL.RequestURI()
  )

  app.logger.Error(err.Error(), "method", method, "url", url)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
  // generic helper for sending JSON formatted error messages to the client with a given status code.
  env := envelope{"error": message}

  err := app.writeJSON(w, status, env, nil)
  if err != nil {
    app.logError(r, err)
    w.WriteHeader(500)
  }
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
  // used when our application encounters an unexpected problem at runtime. Logs detailed error message, then uses errorResponse() helper to send 500 Internal Server Error to client
  app.logError(r, err)

  message := "the server encountered a problem and could not process your request"
  app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
  // method will be used to send a 404 Not Found status code
  message := "the requested resource could not be found"
  app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse (w http.ResponseWriter, r *http.Request) {
  // method will be used to send a 405 Method Not Allowed status code
  message := fmt.Sprintf("this %s method is not supported for this resource", r.Method)
  app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}



