package models

import "gorm.io/gorm"

type PackageRequest {
	gorm.Model
	Name             string  `json:"name"`
	Sender           string  `json:"sender"`
	Receiver         string  `json:"receiver"`
	Senderlocation   string  `json:"sender_location"`
	Receiverlocation string  `jaon:"receiver_location"`
	Fee              int     `json:"fee"`
	Weight           float64 `json:"weight"`
}