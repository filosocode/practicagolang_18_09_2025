package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "postgres://postgres:postgres@localhost:5432/enerbit?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error conectando a PostgreSQL:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("No se pudo hacer ping a PostgreSQL:", err)
	}

	log.Println("Conectado a PostgreSQL")
	DB = db
}
