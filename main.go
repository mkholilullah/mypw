package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roamercodes/mypw/config"
	"github.com/roamercodes/mypw/handler"
	"github.com/roamercodes/mypw/repository"
	"github.com/roamercodes/mypw/usecase"
)

func main() {
	db := config.Connect()

	// inisialisasi repository dan usecase
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// inisialisasi HTTP Handler
	userHandler := handler.NewUserHandler(userUsecase)

	// inisialisasi router
	r := mux.NewRouter()

	// Routing HTTP
	r.HandleFunc("/user/{id}", userHandler.GetUserById).Methods("GET")
	r.HandleFunc("/user/register", userHandler.CreateUser).Methods("POST")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}