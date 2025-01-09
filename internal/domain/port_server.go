package domain

import "encoding/json"

type PortService interface {
	GetUserLogin(*User) (*string, *UserErro)
	GetUser(*string) (*json.RawMessage, *UserErro)
	DeleteUser(input *InputUser) (*json.RawMessage, *UserErro)
	GetUserOne(input *InputUser) (*json.RawMessage, *UserErro)
	SetUser(input *InputUser) (*json.RawMessage, *UserErro)
}
