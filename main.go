package main

import (
	"log"
	"net/http"

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

	// Routing HTTP
	http.HandleFunc("/users", userHandler.GetUserById)
	http.HandleFunc("/user/create", userHandler.CreateUser)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}