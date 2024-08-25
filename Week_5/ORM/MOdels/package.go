package models

import (
	"log"
	"orm/database/models"

	"gorm.io/gorm"
)

type Package struct {
	gorm.Model
	Name             string  `json:"name"`
	Sender           string  `json:"sender"`
	Receiver         string  `json:"receiver"`
	Senderlocation   string  `json:"sender_location"`
	Receiverlocation string  `jaon:"receiver_location"`
	Fee              int     `json:"fee"`
	Weight           float64 `json:"weight"`
}

func MigrsteDB() {
	err := DB.AutoMigrate(&models.Package{})
	if err != nil {
		log.Fatal("Migration Failed", err)
	}
	log.Println("Migration Success")

}
