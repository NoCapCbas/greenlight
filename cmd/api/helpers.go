package main

import (
  "errors"
  "encoding/json"
  "net/http"
  "strconv"
  "fmt"
  "io"
  "strings"

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

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) (error) {

  // limits the size of the request body to 1MB to prevent dos attacks
  maxBytes := 1_048_576
  r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
  // initialize decoder and call DisallowUnkownFields to prevent fields that can not be mapped to movie
  // this will return an error if there is an unkown field
  dec := json.NewDecoder(r.Body)
  dec.DisallowUnknownFields()

  // Decode the request body to the destination 
  err := dec.Decode(dst)
  if err != nil {
    var syntaxError *json.SyntaxError
    var unmarshalTypeError *json.UnmarshalTypeError
    var invalidUnmarshalError *json.InvalidUnmarshalError
    var maxBytesError *http.MaxBytesError

    switch {
      case errors.As(err, &syntaxError):
        return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

      case errors.Is(err, io.ErrUnexpectedEOF):
        return errors.New("body contains badly formed JSON")

      case errors.As(err, &unmarshalTypeError):
        if unmarshalTypeError.Field != "" {
          return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
        }
        return fmt.Errorf("body contains incorrect JSON (at character %d)", unmarshalTypeError.Offset)
      // request body must have data
      case errors.Is(err, io.EOF):
        return errors.New("body must not be empty")
      // JSON field which should not exist
      case strings.HasPrefix(err.Error(), "json: unknown field "):
        fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
        return fmt.Errorf("body contains unknown key %s", fieldName)

      // prevents dos attacks, by limiting request body size 
      case errors.As(err, &maxBytesError): 
        return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

      case errors.As(err, &invalidUnmarshalError):
        panic(err)

      default:
        return err
    }
  }
  
  // using a pointer to an empty struct as the destination, if the is additional JSON data 
  // return error
  err = dec.Decode(&struct{}{})
  if !errors.Is(err, io.EOF) {
    return errors.New("body must only contain a single JSON value")
  }


  return nil
}
