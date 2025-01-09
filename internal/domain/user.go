package domain

import "encoding/json"

type User struct {
	Data json.RawMessage `json:"data"`
}

type UserErro struct {
	Err       error
	CodStatus int
	Body      string
}

type InputUser struct {
	Token     string
	TokenUser string
	Body      json.RawMessage
}
