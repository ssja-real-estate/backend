package util

import (
	"strings"
)

type JError struct {
	Error string `json:"error"`
}

var PageLen = 20

func NewJError(err error) JError {
	jerr := JError{

		Error: "generic error"}
	if err != nil {
		jerr.Error = err.Error()
	} 
	return jerr
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
