package model

import (
	"fmt"
)

type Model struct {
	archive
	comment
	user
	ok        bool
	modelType string
	message   string
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
