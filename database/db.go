package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DatabaseServer struct {
	DB *sql.DB
}

func NewDatabaseServer() *DatabaseServer {
	connectionString := "postgres://postgres:postgres@localhost:5432/gobanks_db?sslmode=disable"
	db, oErr := sql.Open("postgres", connectionString)
	if oErr != nil {
		log.Fatal("Database Unable To Connect")
	}
	if pErr := db.Ping(); pErr != nil {
		log.Fatal("Database Unable To Connect")
	}
	return &DatabaseServer{DB: db}
}

func (db *DatabaseServer) Setup() {
	db.createAccountTable()
	db.createSessionTable()
}

func (db *DatabaseServer) createAccountTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS "_accounts" (
			"id" SERIAL PRIMARY KEY,
			"first_name" VARCHAR(255),
			"last_name" VARCHAR(255),
			"number" SERIAL,
			"balance" FLOAT,
			"created_at" timestamp
		);
	`

	if _, err := db.DB.Exec(sql); err != nil {
		log.Fatalln("Unable to create table in database", err)
	}
}

func (db *DatabaseServer) createSessionTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS "_sessions" (
			"user_id" INT,
			"access_token" VARCHAR(255),
			PRIMARY KEY ("user_id", "access_token")
		);
	`

	if _, err := db.DB.Exec(sql); err != nil {
		log.Fatalln("Unable to create table in database", err)
	}
}
