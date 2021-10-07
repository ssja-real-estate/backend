package util

import (
	"strings"
)

type JError struct {
	ErrorCode int    `json:"error_code"`
	Error     string `json:"error"`
}

type ResultOk struct {
	Data interface{} `json:"data"`
}

func NewJError(err error) JError {
	jerr := JError{
		ErrorCode: 0,
		Error:     "generic error"}
	if err != nil {
		jerr.Error = err.Error()
	}
	return jerr
}

func NewRresult(status int, result interface{}) ResultOk {
	resultok := ResultOk{
		Data: result,
	}
	return resultok
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
