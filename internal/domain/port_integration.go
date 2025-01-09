package domain

import "encoding/json"

type PortIntegration interface {
	GetUserLoginIntegration(*User) (string, *UserErro)
	GetUserIntegration(*string) (*json.RawMessage, *UserErro)
	DeleteUserIntegration(*InputUser) (*json.RawMessage, *UserErro)
	GetUserOneIntegration(*InputUser) (*json.RawMessage, *UserErro)
	SetUserIntegration(*InputUser) (*json.RawMessage, *UserErro)
}
