package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabasePostgres(conf *Config) *gorm.DB {
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: time.Microsecond, 
	// 		LogLevel: logger.Info,
	// 		IgnoreRecordNotFoundError: true,
	// 		ParameterizedQueries: true,
	// 		Colorful: false,
	// 	},
	// )

	DB, err := gorm.Open(postgres.Open(conf.DatabaseURL), &gorm.Config{
		// Logger: newLogger,
	})

	if err != nil {
		panic("Connection failed!")
	}

	return DB
}