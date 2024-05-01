package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DataBase struct {
	Conn *sqlx.DB
}

func New() DataBase {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	urlStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, database, port)

	db, err := sqlx.Connect("postgres", urlStr)

	if err != nil {
		log.Fatal("Could not connect to database --- ", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Could not connect to database --- ", pingErr)
	}

	return DataBase{Conn: db}
}
