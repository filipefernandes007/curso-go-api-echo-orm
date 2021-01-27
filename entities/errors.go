package entities

import "errors"

var ErrUserDoesNotExist = errors.New("user does not exist")

type OutputError struct {
	Error string `json:"error"`
}