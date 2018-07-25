package model

import (
	"github.com/jinzhu/gorm"
	"blog/global"
	"time"
)

type Article struct {
	Object `json:"object,omitempty"`
	gorm.Model
	Title string `gorm:"not null"`
	Content string `gorm:"type:text"`
	Author string `gorm:"not null"`
}

func FindArticleById(id uint) *Article {
	db := global.GLOBAL.DB
	var a Article
	if db.Find(&a, id).RecordNotFound() {
		return nil
	}
	return &a
}

func GetAllArticles() *[]Article {
	var a []Article
	db := global.GLOBAL.DB
	db.Find(&a)
	return &a
}

func AddArticle(article Article) error {
	db := global.GLOBAL.DB
	return db.Create(&article).Error
}

func UpdateArticle(article Article) error {
	db := global.GLOBAL.DB
	return db.Model(&article).Updates(article).Error
}

func DeleteArticle(id uint) error {
	db := global.GLOBAL.DB
	a := Article{Model: gorm.Model{
		ID:        id,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}}
	return db.Delete(&a).Error
}