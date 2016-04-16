package main

// Farid Abulias

import (
	"database/sql"
	"log"
	//"os"
	"fmt"
    _ "github.com/lib/pq"
)

var db *sql.DB = nil
var err error

func connect_db() {

	db, err = sql.Open("postgres","postgres://tbllgrkjejpwzv:e3D-VEc5BmjTyw6pESuJnzgQAo@ec2-54-221-249-201.compute-1.amazonaws.com:5432/dcvc2lb7meb7j5")
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
	 
	  
	rut := "123"
	pass := "123"
	//sql=()
	
	var nombre string
	
	aux :=db.QueryRow("SELECT nombre FROM Cliente WHERE rut=$1 AND password=$2",rut,pass).Scan(&nombre)
	switch {
	case aux==sql.ErrNoRows:
		log.Printf("No user")
	case aux!=nil :
		log.Fatal(aux)
	default:
		fmt.Printf("Nombre %s", nombre)
	}

}