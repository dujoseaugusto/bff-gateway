package integration

import (
	"bff-gatway/internal/domain"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type UserIntegration struct{}

func (d UserIntegration) GetUserLoginIntegration(input *domain.User) (string, *domain.UserErro) {

	var erroReponse domain.UserErro
	baseURL := gethost()["auth"]
	client := resty.New()

	url := fmt.Sprintf("%s/v1/login", baseURL)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(input.Data).
		Post(url)

	if err != nil {
		erroReponse = domain.UserErro{
			Err:       err,
			CodStatus: http.StatusInternalServerError,
		}
		return "", &erroReponse
	}

	if resp.StatusCode() != http.StatusOK {
		erroReponse.Body = string(resp.Body())
		erroReponse.CodStatus = resp.StatusCode()
		erroReponse.Err = fmt.Errorf("Erro in integration login user")
		return "", &erroReponse
	}

	return string(resp.Body()), nil

}

func (d UserIntegration) GetUserIntegration(token *string) (*json.RawMessage, *domain.UserErro) {
	var erroReponse domain.UserErro
	baseURL := gethost()["auth"]
	client := resty.New()

	url := fmt.Sprintf("%s/v1/users", baseURL)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(*token).
		Get(url)

	if err != nil {
		erroReponse = domain.UserErro{
			Err:       err,
			CodStatus: http.StatusInternalServerError,
		}
		return nil, &erroReponse
	}

	if resp.StatusCode() != http.StatusOK {
		erroReponse.Body = string(resp.Body())
		erroReponse.CodStatus = resp.StatusCode()
		erroReponse.Err = fmt.Errorf("Erro in integration login user")
		return nil, &erroReponse
	}

	var rawMessage json.RawMessage
	if err := json.Unmarshal(resp.Body(), &rawMessage); err != nil {
		erroReponse = domain.UserErro{
			Err:       err,
			CodStatus: http.StatusInternalServerError,
			Body:      string(resp.Body()),
		}
		return nil, &erroReponse
	}

	return &rawMessage, nil

}

func (d UserIntegration) DeleteUserIntegration(input *domain.InputUser) (*json.RawMessage, *domain.UserErro) {
	var erroReponse domain.UserErro
	baseURL := gethost()["auth"]
	client := resty.New()

	url := fmt.Sprintf("%s/v1/users/%s", baseURL, input.TokenUser)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(*&input.Token).
		Delete(url)

	if err != nil {
		erroReponse = domain.UserErro{
			Err:       err,
			CodStatus: http.StatusInternalServerError,
		}
		return nil, &erroReponse
	}

	if resp.StatusCode() != http.StatusOK {
		erroReponse.Body = string(resp.Body())
		erroReponse.CodStatus = resp.StatusCode()
		erroReponse.Err = fmt.Errorf("Erro in integration login user")
		return nil, &erroReponse
	}

	var rawMessage json.RawMessage
	if err := json.Unmarshal(resp.Body(), &rawMessage); err != nil {
		erroReponse = domain.UserErro{
			Err:       err,
			CodStatus: http.StatusInternalServerError,
			Body:      string(resp.Body()),
		}
		return nil, &erroReponse
	}

	return &rawMessage, nil

}

func (d UserIntegration) GetUserOneIntegration(input *domain.InputUser) (*json.RawMessage, *domain.UserErro) {
	var erroReponse domain.UserErro
	baseURL := gethost()["auth"]
	client := resty.New()

	url := fmt.Sprintf("%s/v1/users/%s", baseURL, input.TokenUser)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(*&input.Token).
		Get(url)

	if err != nil {
		erroReponse = domain.UserErro{
			Err:       err,
			CodStatus: http.StatusInternalServerError,
		}
		return nil, &erroReponse
	}

	if resp.StatusCode() != http.StatusOK {
		erroReponse.Body = string(resp.Body())
		erroReponse.CodStatus = resp.StatusCode()
		erroReponse.Err = fmt.Errorf("Erro in integration login user")
		return nil, &erroReponse
	}

	var rawMessage json.RawMessage
	if err := json.Unmarshal(resp.Body(), &rawMessage); err != nil {
		erroReponse = domain.UserErro{
			Err:       err,
			CodStatus: http.StatusInternalServerError,
			Body:      string(resp.Body()),
		}
		return nil, &erroReponse
	}

	return &rawMessage, nil

}

func (d UserIntegration) SetUserIntegration(input *domain.InputUser) (*json.RawMessage, *domain.UserErro) {
	var erroReponse domain.UserErro
	baseURL := gethost()["auth"]
	client := resty.New()

	url := fmt.Sprintf("%s/v1/users/%s", baseURL, input.TokenUser)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(*&input.Token).
		SetBody(input.Body).
		Put(url)

	if err != nil {
		erroReponse = domain.UserErro{
			Err:       err,
			CodStatus: http.StatusInternalServerError,
		}
		return nil, &erroReponse
	}

	if resp.StatusCode() != http.StatusOK {
		erroReponse.Body = string(resp.Body())
		erroReponse.CodStatus = resp.StatusCode()
		erroReponse.Err = fmt.Errorf("Erro in integration login user")
		return nil, &erroReponse
	}

	var rawMessage json.RawMessage
	if err := json.Unmarshal(resp.Body(), &rawMessage); err != nil {
		erroReponse = domain.UserErro{
			Err:       err,
			CodStatus: http.StatusInternalServerError,
			Body:      string(resp.Body()),
		}
		return nil, &erroReponse
	}

	return &rawMessage, nil

}
