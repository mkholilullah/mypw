package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roamercodes/mypw/handler"
	"github.com/roamercodes/mypw/middleware"
	"github.com/roamercodes/mypw/repository"
	"github.com/roamercodes/mypw/usecase"
	"gorm.io/gorm"
)

func routeInit(conn *gorm.DB) *mux.Router {

	// DI
	userRepo := repository.NewUserRepository(conn)

	userUsecase := usecase.NewUserUsecase(userRepo)

	// inisialisasi HTTP Handler
	userHandler := handler.NewUserHandler(userUsecase)

	r := mux.NewRouter()

	r.Use(middleware.LoggerMiddleware)

	r.HandleFunc("/register", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)

	r.HandleFunc("/user/{id}", userHandler.GetUserById).Methods("GET")

	return r

}