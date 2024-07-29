package main

import (
  "errors"
  "encoding/json"
  "net/http"
  "strconv"

  "github.com/julienschmidt/httprouter"
)

type envelope map[string]any

func (app *application) readIDParam(r *http.Request) (int64, error) {
  // ParamsFromContext retrieves a slice of params 
  // from request context 
  params := httprouter.ParamsFromContext(r.Context())
  // get the id param and convert it into a base 10 int
  // if id is less than 1 we know the id is invalid
  id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
  if err != nil || id < 1 {
    return 0, errors.New("invalid id parameter")
  }

  return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) (error) {

  jsonData, err := json.MarshalIndent(data, "", "\t")
  if err != nil {
    return err
  }
  // only for terminal applications fro formatting
  jsonData = append(jsonData, '\n')
  // lopp through headers to set
  for key, value := range headers {
    w.Header()[key] = value
  }
  
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  w.Write(jsonData)

  return nil
}
