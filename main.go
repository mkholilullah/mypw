package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/roamercodes/mypw/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(" Error loading .env files : %v", err)
	}

	conf := config.GetConfig()
	conn := config.InitDatabasePostgres(&conf)
	config.AutoMigrate(conn)

	fmt.Printf("Server starting at localhost:%v ...\n", conf.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", conf.Port), routeInit(conn))
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}

	// db := config.Connect()
	// userRepo := repository.NewUserRepository(db)
	// userUsecase := usecase.NewUserUsecase(userRepo)

	// userHandler := handler.NewUserHandler(userUsecase)

	// r := mux.NewRouter()

	// r.HandleFunc("/user/{id}", userHandler.GetUserById).Methods("GET")
	// r.HandleFunc("/user/register", userHandler.CreateUser).Methods("POST")
	// r.HandleFunc("/user/login", userHandler.Login).Methods("POST")
	// log.Println("Server running on :8080")
	// log.Fatal(http.ListenAndServe(":8080", r))
}