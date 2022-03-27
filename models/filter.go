package models

type Filter struct {
	Header HeadFilter `json:"header"`
	Form   Form       `json:"form"`
}
