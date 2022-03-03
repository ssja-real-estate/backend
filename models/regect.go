package models

import "time"

type Reject struct {
	RejectDate  time.Time `json:"rejectDate"`
	Description string    `json:"description"`
	Rejected    bool      `json:"rejected"`
}
