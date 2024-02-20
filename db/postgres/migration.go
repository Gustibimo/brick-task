package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"os"
)

func Init() {
	db := DbConfig()
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	// Set up the database dialect
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	// Run the migrations
	if err := goose.Up(db, "db/postgres"); err != nil {
		panic(err)
	}

	fmt.Println("Migrations applied successfully.")
}

func DbConfig() *sql.DB {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		panic(err)
	}

	return db
}
