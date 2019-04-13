package utils

import (
    "github.com/google/uuid"
)


// UUID should really be made more general, but this was so easy
// It generates a 128bit unique identifier, stored as a string
// https://en.wikipedia.org/wiki/Universally_unique_identifier
func UUID() string {
  return uuid.New().String()
}
