package config

import (
	"fmt"
	"log"
	"os"

	"github.com/roamercodes/mypw/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbUsername = "DB_USERNAME"
	dbPassword = "DB_PASSWORD"
	dbHost 	   = "DB_HOST"
	dbPort     = "DB_PORT"
	dbName     = "DB_NAME"
	serverPort = "SERVER_PORT"
)

type Config struct {
	Port string
	DatabaseURL string
}

func GetConfig() Config {
	return Config{
		Port: os.Getenv(serverPort),
		DatabaseURL: fmt.Sprintf("postgres://%v:%v@%v:%v/%v", os.Getenv(dbUsername), os.Getenv(dbPassword), os.Getenv(dbHost),
					 os.Getenv(dbPort), os.Getenv(dbName)), 
	}
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{},
	)

	if err != nil {
		log.Fatalf("Migration failed . Error : %v", err)
	}
	log.Println("Migration Success....")
}

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password='' dbname=db_mypass port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db
}
