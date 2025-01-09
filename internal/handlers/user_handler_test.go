package handlers_test

// import (
// 	"bff-gatway/internal/domain"
// 	"bff-gatway/internal/handlers"
// 	"bff-gatway/internal/helpers"
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// type MockUserService struct {
// }

// func (d MockUserService) GetUser(input *domain.User) (*json.RawMessage, error) {
// 	response := json.RawMessage(`{
// 		"status": "success",
// 		"data": {
// 			"original": {
// 				"type": "application/vnd.lime.collection+json",
// 				"itemType": "application/vnd.lime.document-select+json",
// 				"items": [
// 					{
// 						"header": {
// 							"type": "application/vnd.lime.media-link+json",
// 							"value": {
// 								"title": "[1] KKMS TUR",
// 								"text": "\n"
// 							}
// 						},
// 						"options": [
// 							{
// 								"label": {
// 									"type": "text/plain",
// 									"value": "Selecionar"
// 								},
// 								"value": {
// 									"type": "text/plain",
// 									"value": 1
// 								}
// 							}
// 						],
// 						"extras": []
// 					}
// 				],
// 				"extras": {
// 					"data": {
// 						"id": 9388,
// 						"title": "KKMS TUR",
// 						"tags": [
// 							"id:2602",
// 							"razaoSocial: KAIQUE MARQUES SILVA 02130961290",
// 							"cnpj:27.054.919/0001-14",
// 							"cidade: MANAUS",
// 							"uf:AM",
// 							"endereço: RUA PROFESSORA LEA ALENCAR, 322",
// 							"bairro: ALVORADA",
// 							"cep:69042-050",
// 							"segmentacao:FRANQUEADO",
// 							"coordenadas: -3.077, -60.04819"
// 						]
// 					}
// 				},
// 				"selected": false,
// 				"details": true
// 			}
// 		}
// 	}`)
// 	return &response, nil
// }

// type MockUserServiceError struct {
// }

// func (d MockUserServiceError) GetUser(input *domain.User) (*json.RawMessage, error) {
// 	return nil, errors.New("Error")
// }

// var mockLogger = helpers.NewLogger(log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))

// func TestHandleGetUser(t *testing.T) {
// 	userRequest := domain.User{
// 		Data: json.RawMessage(`{"teamToken": "35f7c2a6-7039-41ae-9f88-85e7f7b455cf", "botToken": "c495ec47-6986-42c7-a73e-f7f2f543ea44", "option": "KKMS TUR", "showNumbers": true, "showMoreTitle": "Ver mais", "currentPage": "1", "limit": "2"}`),
// 	}

// 	mockService := &MockUserService{}

// 	handler := handlers.NewHandler(mockService, mockLogger)

// 	reqBody, _ := json.Marshal(userRequest)

// 	req, _ := http.NewRequest("POST", "/user-content/829/15", bytes.NewBuffer(reqBody))

// 	rec := httptest.NewRecorder()

// 	handler.HandleGetUserLogin(rec, req)

// 	assert.Equal(t, rec.Code, http.StatusOK)

// }

// func TestHandleGetUserError(t *testing.T) {
// 	userRequest := domain.User{
// 		Data: json.RawMessage(`{"teamToken": "35f7c2a6-7039-41ae-9f88-85e7f7b455cf", "botToken": "c495ec47-6986-42c7-a73e-f7f2f543ea44", "option": "KKMS TUR", "showNumbers": true, "showMoreTitle": "Ver mais", "currentPage": "1", "limit": "2"}`),
// 	}

// 	mockService := &MockUserServiceError{}

// 	handler := handlers.NewHandler(mockService, mockLogger)

// 	reqBody, _ := json.Marshal(userRequest)

// 	req, _ := http.NewRequest("POST", "/user-content/829/15", bytes.NewBuffer(reqBody))

// 	rec := httptest.NewRecorder()

// 	handler.HandleGetUserLogin(rec, req)

// 	assert.Equal(t, rec.Code, http.StatusInternalServerError)

// }
