package main

import (
  // "fmt"
  "net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
  // map that holds data to be sent to client
  env := envelope{
    "status": "available",
    "system_info": map[string]string{
      "environment": app.config.env,
      "version": version,
    },
  }

  // pass map th json marshal
  err := app.writeJSON(w, http.StatusOK, env, nil)
  if err != nil {
    app.serverErrorResponse(w, r, err)
  }
    
}
