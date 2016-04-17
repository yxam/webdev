package main

// Farid Abulias

import (
	"database/sql"
	"log"
	"os"

)

var db *sql.DB = nil
var err error

type information struct {
	rut  string //`form:"rut"`// json:"rut" binding:"required"
	pass string //`form:"pass"`// json:"pass" binding:"required"`
}

type account_s struct {
	rut string
	saldo int
	tipo int
}

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
	create[1] = "CREATE TABLE IF NOT EXISTS Banco (id serial, nombre varchar(50) NOT NULL, PRIMARY KEY (id))"
	//serial = (int) auto_increment
    create[2] = "CREATE TABLE IF NOT EXISTS Cuenta(id bigint, rut_cliente varchar(12) REFERENCES cliente(rut), tipo integer NOT NULL, saldo integer NOT NULL"
    //id = numero de cuenta, por eso bigint y no serial que es auto incremental. 
    create[3] ="CREATE TABLE IF NOT EXISTS Transferencia(rut_origen varchar(12) REFERENCES cliente(rut), rut_destino varchar(12) NOT NULL,monto integer NOT NULL, fecha timestamp,PRIMARY KEY (rut_origen,fecha))" 
    //timestamp, guarda fecha y hora
    var length=cap(create)
    i:=0
    for i<length { 
	    _, err := db.Exec(create[i])    
	    if err != nil {
	        log.Fatalf("Error creating table: %q", err)
	    }
	 i = i+1 
	}
	disconnect_db()
}

func login(client information) bool {
	connect_db()
	var nombre string
	//sql=("SELECT * FROM Cliente WHERE rut=? AND pass=?") Nose si funciona pasandole aqui las variables
	//var aux  para verificar si encontro o no a la persona, ya que si ejecuta la query aunque no encuentre nada retornara TRUE.
	err := db.QueryRow("SELECT nombre FROM Cliente WHERE rut=? AND pass=?", client).Scan(&client)
    switch {
	    case err == sql.ErrNoRows:
	    	 disconnect_db()
	    	 return false
	    case err != nil:
	         disconnect_db()
	         return false
	    default:
	    	 disconnect_db()
	    	 return true
    }
 
}

func account(client information) account_s {
	connect_db()
	row := db.QueryRow("SELECT * FROM Cuenta WHERE Cuenta.rut == ?", client.rut).Scan(&client.rut)
	switch {
		case row == sql.ErrNoRows:	
			disconnect_db()
			return nil
		case row != nil:
			disconnect_db()
			return nil
		default:
			var tmp account_s
			tmp.rut = row.rut
			tmp.saldo = row.saldo
			tmp.tipo = row.tipo
			disconnect_db()
			return tmp
			
	}
}
