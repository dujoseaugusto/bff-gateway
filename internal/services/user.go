package services

import (
	"bff-gatway/internal/domain"
	"bff-gatway/internal/integration"
	"encoding/json"
)

type UserService struct {
	Integratio domain.PortIntegration
}

func NewUserService() *UserService {

	return &UserService{
		Integratio: &integration.UserIntegration{},
	}
}

func (d UserService) GetUserLogin(input *domain.User) (*string, *domain.UserErro) {
	list, err := d.Integratio.GetUserLoginIntegration(input)
	return &list, err
}

func (d UserService) GetUser(token *string) (*json.RawMessage, *domain.UserErro) {
	return d.Integratio.GetUserIntegration(token)
}

func (d UserService) DeleteUser(input *domain.InputUser) (*json.RawMessage, *domain.UserErro) {
	return d.Integratio.DeleteUserIntegration(input)
}

func (d UserService) GetUserOne(input *domain.InputUser) (*json.RawMessage, *domain.UserErro) {
	return d.Integratio.GetUserOneIntegration(input)
}

func (d UserService) SetUser(input *domain.InputUser) (*json.RawMessage, *domain.UserErro) {
	return d.Integratio.SetUserIntegration(input)
}
