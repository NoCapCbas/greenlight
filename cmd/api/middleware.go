package main

import (
  "fmt"
  "net/http"
)

func (app *application) recoverPanic(next http.Handler) (http.Handler) {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // creates defered function which will always be run in the event of a panic
    defer func() {
      // uses builtin recover function to check if there has been a panic
      if err := recover(); err != nil {
        // closes current connection after a response has been sent
        w.Header().Set("connection", "close")
        // recover() returns type any so we format it into type error and log error use helper function 
        app.serverErrorResponse(w, r, fmt.Error("%s", err))
      }
    }()
    next.ServeHttp(w, r)
  })
}
