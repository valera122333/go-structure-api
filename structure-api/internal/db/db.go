package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	dsn := os.Getenv("DB_DSN")

	var database *gorm.DB
	var err error

	// ждём пока postgres запустится
	for i := 0; i < 10; i++ {

		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			log.Println("connected to database")
			break
		}

		log.Println("waiting for database...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("failed to connect db:", err)
	}

	runMigrations(dsn)

	return database
}

func runMigrations(dsn string) {

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = goose.Up(sqlDB, "migrations")
	if err != nil {
		log.Fatal("migration failed:", err)
	}

	log.Println("migrations applied")
}
