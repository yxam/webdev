package modelutil

import (
	"database/sql"
	"log"
	//"os"

	_ "github.com/lib/pq"
)

var db *sql.DB = nil
var err error

type Information struct {
	Rut  string 
	Pass string //`form:"pass"`// json:"pass" binding:"required"`
}

type account_s struct {
	rut string
	saldo int
	tipo int
}

func connect_db() {
	db, err = sql.Open("postgres", "postgres://tbllgrkjejpwzv:e3D-VEc5BmjTyw6pESuJnzgQAo@ec2-54-221-249-201.compute-1.amazonaws.com:5432/dcvc2lb7meb7j5")
		//os.Getenv("DATABASE_URL"))
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

func Init() bool {
	connect_db()
    
    var create [4]string
	create[0] = "CREATE TABLE IF NOT EXISTS Cliente (rut varchar(12), pass varchar(4) NOT NULL,	PRIMARY KEY(rut))"
	create[1] = "CREATE TABLE IF NOT EXISTS Banco (id serial, nombre varchar(50) NOT NULL, PRIMARY KEY (id))"
	create[2] = "CREATE TABLE IF NOT EXISTS Cuenta(id bigint, rut_cliente varchar(12) REFERENCES cliente(rut), tipo integer NOT NULL, saldo integer NOT NULL)"
    create[3] ="CREATE TABLE IF NOT EXISTS Transferencia(rut_origen varchar(12) REFERENCES cliente(rut), rut_destino varchar(12) NOT NULL,monto integer NOT NULL, fecha timestamp,PRIMARY KEY (rut_origen,fecha))" 
    var length = cap(create)
    i := 0
    for i < length { 
	    _, err := db.Exec(create[i])    
	    if err != nil {
	        disconnect_db()
	        return false
	    }
	    i++
	}
	disconnect_db()
	return true
}

func Login(rut, pass string) bool {
	connect_db()
	_, err := db.Query("SELECT * FROM Cliente WHERE rut=$1 AND pass=$2", rut, pass)
    disconnect_db()
    if err != nil {
    	return false
    } else {
    	return true
    }
}

func Account(rut string) *account_s {
	connect_db()
	tmp := new(account_s) 
	row := db.QueryRow("SELECT * FROM Cuenta WHERE Cuenta.rut == ?", rut).Scan(&tmp)
	disconnect_db()
	switch {
		case row == sql.ErrNoRows:	
			return nil
		case row != nil:
			return nil
		default:
			return tmp
	}
}
