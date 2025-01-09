package routes

import (
	"bff-gatway/internal/handlers"
	"bff-gatway/internal/helpers"
	middleware "bff-gatway/internal/helpers"
	"bff-gatway/internal/services"

	_ "bff-gatway/docs/swagger"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(logger *helpers.Logger) *mux.Router {
	service := services.NewUserService()
	handler := handlers.NewHandler(service, logger)

	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.Use(middleware.RequestID)
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", handler.HandleGetUserLogin).Methods("POST")
	authRouter.HandleFunc("/users", handler.HandleGetUsers).Methods("GET")
	authRouter.HandleFunc("/users", handler.HandleGetUserRegister).Methods("PUT")
	authRouter.HandleFunc("/users/{token}", handler.HandleGetUser).Methods("GET")
	authRouter.HandleFunc("/users/{token}", handler.HandleGetUserRemove).Methods("DELETE")

	return router
}
