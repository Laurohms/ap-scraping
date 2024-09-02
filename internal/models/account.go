package models

import "time"

type SubAccount string

type Postings struct {
	SubAccount  SubAccount `json:"sub_account,omitempty"`
	Date        time.Time  `json:"date,omitempty"`
	Description string     `json:"description,omitempty"`
	Subdivision string     `json:"subdivision,omitempty"`
	Value       float64    `json:"value,omitempty"`
	Balance     float64    `json:"balance,omitempty"`
}
