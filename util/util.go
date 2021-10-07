package util

import (
	"strings"
)

type JError struct {
	Error string `json:"error"`
}

type ResultOk struct {
	Data interface{} `json:"data"`
}

func NewJError(err error) JError {
	jerr := JError{

		Error: "generic error"}
	if err != nil {
		jerr.Error = err.Error()
	}
	return jerr
}

func NewRresult(result interface{}) ResultOk {
	resultok := ResultOk{
		Data: result,
	}
	return resultok
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
