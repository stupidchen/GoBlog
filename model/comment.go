package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	Object
	gorm.Model
	Author string `gorm:"not null"`
	Content string `gorm:"type:text"`
}
