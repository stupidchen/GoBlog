package view

import (
	"blog/model"
	"encoding/json"
	"blog/global"
)

type RequestData struct {
	Article	model.Article `json:"article"`
	Comment	model.Comment `json:"comment"`
	User	model.User	  `json:"user"`
}

type ResponseData struct {
	Ok        bool
	ModelType string
	Message   string
	Object    model.Object
}

func (m *ResponseData) ToString() string {
	b, err := json.Marshal(m)
	if err != nil {
		global.GLOBAL.Logger.Panicf("Cannot marshall the data due to %s.", err.Error())
		return "Cannot format return data"
	}
	return string(b)
}

func FromString(s *string) *RequestData {
	var data RequestData
	err := json.Unmarshal([]byte(*s), &data)
	if err != nil {
		global.GLOBAL.Logger.Panicf("Cannot unmarshall the request body due to %s.", err.Error())
		return nil
	}
	return &data
}

func InitHint(msg string) *ResponseData {
	return &ResponseData{
		Ok:        true,
		ModelType: "Hint",
		Message:   msg,
	}
}

func InitError(msg string) *ResponseData {
	return &ResponseData{
		Ok:        false,
		ModelType: "Error",
		Message:   msg,
	}
}
