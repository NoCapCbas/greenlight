package validator

import (
  "regexp"
  "slices"
)

// regular expression for email addresses
var (
  EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// new validator type, containing map of validator errors
type Validator struct {
  Errors map[string]string
}

func New() *Validator {
  // creates new validator instance with empty  errors map
  return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
  // returns true if the errors map doesn't contain any entries
  return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {

  // adds error message to map (as long as not duplicate entry)
  if _, exists := v.Errors[key]; !exists {
    v.Errors[key] = message
  }
}

func (v *Validator) Check(ok bool, key, message string) {
  // adds error message to map only if validation check is not 'ok'
  if !ok {
    v.AddError(key, message)
  }
}


func PermittedValue[T comparable](value T, permittedValues ...T) (bool) {
  // returns true if specific value is in a list of permitted values
  return slices.Contains(permittedValues, value)
}

func Matches(value string, rx *regexp.Regexp) (bool) {
  // returns true if a string value matches a specific regexp pattern
  return rx.MatchString(value)
}

func Unique[T comparable](values []T) bool {
  // returns true if all values in a slice are unique
  uniqueValues := make(map[T]bool)

  for _, value := range values {
    uniqueValues[value] = true
  }

  return len(values) == len(uniqueValues)
}







