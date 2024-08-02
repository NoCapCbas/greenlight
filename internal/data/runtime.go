package data 

import (
  "fmt"
  "strconv"
  "errors"
  "strings"
)

// Error type for unsuccessful JSON conversion
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
  // generate a string containing the movie runtime in the required format 
  jsonValue := fmt.Sprintf("%d mins", r)
  
  // Use the strconv.Quote() function on the string to wrap it in double quotes.
  // It needs to be surrounded by double quotes in order to be the valid json string
  quotedJSONValue := strconv.Quote(jsonValue)
  
  // Convert the quoted string value to a byte slice and return it 
  return []byte(quotedJSONValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) (error) {
  // expected jsonValue "<runtime> mins"
  // remote surrounding double-quotes
  unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
  if err != nil {
    // return error if unable to remove double-quotes json value
    return ErrInvalidRuntimeFormat
  }

  // split string to isolate number from text
  parts := strings.Split(unquotedJSONValue, " ")

  // sanity check, make sure string is expected format
  if len(parts) != 2 || parts[1] != "mins" {
    // return error if sanity check fails
    return ErrInvalidRuntimeFormat
  }

  // parse string containing number into int32
  i, err := strconv.ParseInt(parts[0], 10, 32)
  if err != nil {
    // return error if first part of string does not parse to int
    return ErrInvalidRuntimeFormat
  }
  
  // covert int32 to runtime type and assign receiver
  *r = Runtime(i)

  return nil
}

