package main

// Farid Abulias

import (
	"database/sql"
	"log"
	"os"
    "fmt"
    "github.com/lib/pq"
	
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

func init() {
	connect_db()
	var create []string
	create[0] = "CREATE TABLE IF NOT EXISTS Cliente (rut varchar(12), pass varchar(4) NOT NULL,	PRIMARY KEY(rut)"
	create[1] = "CREATE TABLE IF NOT EXISTS Banco (id serial,	nombre varchar(50) NOT NULL, PRIMARY KEY (id))"
	//serial = (int) auto_increment
    create[2] = "CREATE TABLE IF NOT EXISTS Cuenta(id bigint, rut_cliente varchar(12) REFERENCES cliente(rut), tipo text NOT NULL,	saldo integer NOT NULL"
    //id = numero de cuenta, por eso bigint y no serial que es auto incremental. 
    create[3] ="CREATE TABLE IF NOT EXISTS Transferencia(rut_origen varchar(12) REFERENCES cliente(rut), rut_destino varchar(12) NOT NULL,monto integer NOT NULL, fecha timestamp,PRIMARY KEY (rut_origen,fecha))" 
    //timestamp, guarda fecha y hora
    var length=cap(create)
    i:=0
    for i<length{ 
	    _, err := db.Exec(create[i])    
	    if err != nil {
	        log.Fatalf("Error creating table: %q", err)
	    }
	 i = i+1 
	}
    


		

	disconnect_db()
}

func login(rut, pass string) {
	connect_db()
	var nombre string
	//sql=("SELECT * FROM Cliente WHERE rut=? AND pass=?") Nose si funciona pasandole aqui las variables
	//var aux  para verificar si encontro o no a la persona, ya que si ejecuta la query aunque no encuentre nada retornara TRUE.
	aux := db.QueryRow("SELECT nombre FROM Cliente WHERE rut=? AND pass=?",rut,pass).Scan(&nombre) //Sí no es así QueryRow("sql",rut, pass)
    switch {
    case aux==sql.ErrNoRows:
    	 log.Printf("No user")
    case aux!=nil:
         log.Fatal(aux)
    default:
    	 fmt.Printf("User: %s", nombre)
    }
 
}
