package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Object `json:"object,omitempty"`
	gorm.Model
	Username string `gorm:"not null"`
	Email string `gorm:"not null"`
	Password string `gorm:"not null"`
	Info string `gorm:"type:text"`
}
