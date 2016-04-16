package main

// Farid Abulias

import (
	"database/sql"
	"log"
	"os"
	"fmt"
    _ "github.com/bmizerany/pq"
)

var db *sql.DB = nil
var err error

func connect_db() {

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }
}

func disconnect_db() {
	err = db.Close()
	if err != nil {
		log.Fatalf("Error closing database: %q", err)
	}
}

func main() {
	connect_db()
	 
	  
	rut := "1111"
	pass := "111"
	//sql=()
	
	var nombre string
	
	aux :=db.QueryRow("SELECT * FROM Cliente WHERE rut=? AND pass=?",rut,pass).Scan(&nombre)
	switch {
	case aux==sql.ErrNoRows:
		log.Printf("No user")
	case aux!=nil :
		log.Fatal(aux)
	default:
		fmt.Printf("Nombre %s", nombre)
	}

}