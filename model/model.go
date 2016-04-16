package main

// Farid Abulias

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB = nil
var err error

func connect_db() {
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }
    return db
}

func disconnect_db() {
	err = db.Close()
	if err != nil {
		log.Fatalf("Error closing database: %q", err)
	}
}

func init() {
	connect_db()
	client_table := "CREATE TABLE IF NOT EXISTS Cliente (
		rut	varchar(12) NOT NULL,
		pass varchar(4) NOT NULL,
		UNIQUE(rut),
		PRIMARY KEY(rut)"
		
	_, err := db.Exec(client_table)
	if err != nil {
		return log.Fatalf("Error creating table: %q", er)
	}



		

	disconnect_db()
}

func select(rut, pass string) {
	connect_db()
	db.Exec("SELECT rut, pass ")
}
