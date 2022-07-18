package models

import "time"

type SadadPayment struct {
	MerchantId    string    `json:"MerchantId"`
	TerminalId    string    `json:"TerminalId"`
	Amount        int       `json:"Amount"`
	OrderId       int64     `json:"OrderId"`
	LocalDateTime time.Time `json:"LocalDateTime"`
	ReturnUrl     string    `json:"ReturnUrl"`
	SignData      string    `json:"SignData"`
}
