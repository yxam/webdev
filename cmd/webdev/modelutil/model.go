package modelutil
<<<<<<< HEAD:cmd/webdev/modelutil/model.go
=======

// Farid Abulias
>>>>>>> e9c67ed45d7ea15213d21eb05718d6947ffdfe71:cmd/webdev/modelutil/model.go

import (
	"database/sql"
	"log"
	"os"

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

func Init() bool {
<<<<<<< HEAD:cmd/webdev/modelutil/model.go
	connect_db()
=======
	db, err = sql.Open("postgres", "postgres://tbllgrkjejpwzv:e3D-VEc5BmjTyw6pESuJnzgQAo@ec2-54-221-249-201.compute-1.amazonaws.com:5432/dcvc2lb7meb7j5")
    if err != nil {
        return false
    }
	defer db.Close()
    
>>>>>>> e9c67ed45d7ea15213d21eb05718d6947ffdfe71:cmd/webdev/modelutil/model.go
    var create []string
	create[0] = "CREATE TABLE IF NOT EXISTS Cliente (rut varchar(12), pass varchar(4) NOT NULL,	PRIMARY KEY(rut))"
	create[1] = "CREATE TABLE IF NOT EXISTS Banco (id serial, nombre varchar(50) NOT NULL, PRIMARY KEY (id))"
	//serial = (int) auto_increment
    create[2] = "CREATE TABLE IF NOT EXISTS Cuenta(id bigint, rut_cliente varchar(12) REFERENCES cliente(rut), tipo integer NOT NULL, saldo integer NOT NULL)"
    //id = numero de cuenta, por eso bigint y no serial que es auto incremental. 
    create[3] ="CREATE TABLE IF NOT EXISTS Transferencia(rut_origen varchar(12) REFERENCES cliente(rut), rut_destino varchar(12) NOT NULL,monto integer NOT NULL, fecha timestamp,PRIMARY KEY (rut_origen,fecha))" 
    //timestamp, guarda fecha y hora
    var length = cap(create)
    i := 0
    for i < length { 
	    _, err := db.Exec(create[i])    
	    if err != nil {
			//disconnect_db()
<<<<<<< HEAD:cmd/webdev/modelutil/model.go
	        disconnect_db()
=======
>>>>>>> e9c67ed45d7ea15213d21eb05718d6947ffdfe71:cmd/webdev/modelutil/model.go
	        return false
	    }
	    i++
	}
<<<<<<< HEAD:cmd/webdev/modelutil/model.go
	disconnect_db()
=======
	
>>>>>>> e9c67ed45d7ea15213d21eb05718d6947ffdfe71:cmd/webdev/modelutil/model.go
	return true
}

func Login(client Information) bool {
	connect_db()
	tmp := new(Information)
	err := db.QueryRow("SELECT nombre FROM Cliente WHERE rut=? AND pass=?", client).Scan(&tmp)
    switch {
	    case err == sql.ErrNoRows:
	    	 defer disconnect_db()
	    	 return false
	    case err != nil:
	         defer disconnect_db()
	         return false
	    default:
	    	 defer disconnect_db()
	    	 return true
    }
 
}

func Account(rut string) *account_s {
	connect_db()
	tmp := new(account_s) 
	row := db.QueryRow("SELECT * FROM Cuenta WHERE Cuenta.rut == ?", rut).Scan(&tmp)
	switch {
		case row == sql.ErrNoRows:	
			defer disconnect_db()
			return nil
		case row != nil:
			defer disconnect_db()
			return nil
		default:
			return tmp
			
	}
}
