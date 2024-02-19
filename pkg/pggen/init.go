package pggen

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"time"
)

func Init(host string, port string, uname string, pwd string, dbname string, servicename string) *sqlx.DB {

	defer func() {
		if r := recover(); r != nil {
			slog.Error("Errors")
			fmt.Println("Recovered from panic:", r)
		}
	}()

	postgres := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, uname, pwd, dbname)

	db, err := sqlx.Open("postgres", postgres)

	if err != nil {
		slog.Error("Cannot connect to Postgres: ", err)
		panic(err)
	}

	db.SetMaxOpenConns(300)
	db.SetMaxIdleConns(300)
	db.SetConnMaxLifetime(3 * time.Minute)

	return db
}
