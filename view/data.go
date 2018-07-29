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
	Ok        bool		   `json:"ok"`
	ModelType string	   `json:"modelType"`
	Message   string	   `json:"message"`
	Object    model.Object `json:"object,omitempty"`
}

func (m *ResponseData) ToString() string {
	b, err := json.Marshal(m)
	if err != nil {
		global.GLOBAL.Logger.Printf("Cannot marshall the data due to %s.", err.Error())
		return "Cannot marshal return data"
	}
	return string(b)
}

func FromString(s *string) *RequestData {
	var data RequestData
	if len(*s) == 0 {
		global.GLOBAL.Logger.Printf("Cannot unmarshall the empty request body")
		return nil
	}
	err := json.Unmarshal([]byte(*s), &data)
	if err != nil {
		global.GLOBAL.Logger.Printf("Cannot unmarshall the request body due to %s.", err.Error())
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
