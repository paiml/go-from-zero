// Package greet builds a greeting for a name.
package greet

import (
	"errors"
	"strings"
)

// ErrEmptyName is returned when the name is blank.
var ErrEmptyName = errors.New("greet: name is empty")

// Greeting returns "Hello, <name>!" with surrounding space trimmed.
func Greeting(name string) (string, error) {
	n := strings.TrimSpace(name)
	if n == "" {
		return "", ErrEmptyName
	}
	return "Hello, " + n + "!", nil
}
