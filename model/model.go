package model

import (
	"fmt"
	"blog/db"
)

type Model struct {
	Articles
	Comments
	Users
	ok        bool
	modelType string
	message   string
}

func init() {
	db.Db.AutoMigrate(&Article{}, &Comment{}, &User{})
}

func (m *Model) ToString() string {
	return fmt.Sprintf("Model (Type: %s, Content: %s)", m.modelType, m.message)
}

func FromString(s *string) *Model {
	//TODO
	return &Model{
		ok: true,
		modelType: "UNKNOWN",
	}
}

func InitHint(msg string) *Model {
	return &Model{
		ok: true,
		modelType: "Hint",
		message: msg,
	}
}

func InitError(msg string) *Model {
	return &Model{
		ok: false,
		modelType: "Error",
		message: msg,
	}
}
