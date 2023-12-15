package models

import (
	"github.com/jinzhu/gorm"
)


type Message struct {
	gorm.Model

	Sender string `gorm:"not null"`
	Message string `gorm:"not null"`
	Receiver    string `gorm:"not null"`
	
}
