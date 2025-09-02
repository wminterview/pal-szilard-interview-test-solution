package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=" + "localhost" +
		" user=" + "postgres" +
		" password=" + "1234" +
		" dbname=" + "postgres" +
		" port=" + "5432" +
		" sslmode=" + "disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
}
