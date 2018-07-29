package model

import (
	"github.com/jinzhu/gorm"
	"blog/global"
	"time"
)

type Comment struct {
	Object `json:"object,omitempty"`
	gorm.Model
	Article uint `gorm:"not null"`
	Author uint `gorm:"not null"`
	Content string `gorm:"type:text"`
}

func FindCommentById(id uint) *Comment {
	db := global.GLOBAL.DB
	var c Comment
	if db.Find(&c, id).RecordNotFound() {
		return nil
	}
	return &c
}

func FindCommentByArticleId(articleId uint) *[]Comment {
	var comments []Comment
	db := global.GLOBAL.DB
	db.Where("article = ?", articleId).Find(&comments)
	return &comments
}

func AddComment(comment *Comment) error {
	db := global.GLOBAL.DB
	return db.Create(comment).Error
}

func UpdateComment(comment *Comment) error {
	db := global.GLOBAL.DB
	return db.Model(comment).Updates(*comment).Error
}

func DeleteComment(id uint) error {
	db := global.GLOBAL.DB
	a := Comment{Model: gorm.Model{
		ID:        id,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}}
	return db.Delete(&a).Error
}