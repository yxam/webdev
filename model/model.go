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
		rut	varchar(12),
		pass varchar(4) NOT NULL,
		PRIMARY KEY(rut)"
	bank_table := "CREATE TABLE IF NOT EXISTS Banco (
		id serial,
		nombre varchar(50) NOT NULL,
		PRIMARY KEY (id))"
	//serial = (int) auto_increment
    cuenta_table := "CREATE TABLE IF NOT EXISTS Cuenta(
    	id bigint,
    	rut_cliente varchar(12) REFERENCES cliente(rut), 
    	tipo text NOT NULL,
    	saldo integer NOT NULL"
    //id = numero de cuenta, por eso bigint y no serial que es auto incremental. 
    transfer_table:="CREATE TABLE IF NOT EXISTS Transferencia(
    	rut_origen varchar(12) REFERENCES cliente(rut),
    	rut_destino varchar(12) NOT NULL,
    	monto integer NOT NULL,
    	fecha timestamp,
    	PRIMARY KEY (rut_origen,fecha),
    	)" 
    //timestamp, guarda fecha y hora

	_, err := db.Exec(client_table)
    
	if err != nil {
		return log.Fatalf("Error creating table: %q", er)
	}
    


		

	disconnect_db()
}

func login(rut, pass string) {
	connect_db()
	sql=("SELECT * FROM Cliente WHERE rut=? AND pass=?",rut, pass) //Nose si funciona pasandole aqui las variables
	var aux // para verificar si encontro o no a la persona, ya que si ejecuta la query aunque no encuentre nada retornara TRUE.
	_, aux := db.QueryRow(sql) //Sí no es así QueryRow("sql",rut, pass)


}
