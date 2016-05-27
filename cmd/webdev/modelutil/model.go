package modelutil

import (
	"database/sql"
	"log"
	"fmt"
	//"os"
	_ "github.com/lib/pq"
)

var db *sql.DB = nil
var err error

type Information struct {
	Rut  string
	Pass string //`form:"pass"`// json:"pass" binding:"required"`
}

type Account_s struct {
	Id int
	Rut_cliente string
	Tipo int
	Saldo int
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

    var create [6]string
	create[0] = "CREATE TABLE IF NOT EXISTS Cliente(Rut  varchar(12), Password varchar[4],Nombre string, Direccion string, Comuna string, Ciudad string, Telefono string, mail string ,PRIMARY KEY(Rut) )"
	create[1] = "CREATE TABLE IF NOT EXISTS Banco (Id int, Nombre string, Ciudad string, PRIMARY KEY (Id))"
	create[2] = "CREATE TABLE IF NOT EXISTS Cuenta(Nmro_cuenta bigint, rut_cliente varchar(12) REFERENCES Cliente(Rut), Tipo integer REFERENCES Tipo_cuentas(Id) NOT NULL, Saldo integer NOT NULL, PRIMARY KEY(Nmro_cuenta))"
    create[3] = "CREATE TABLE IF NOT EXISTS Transferencia(Id bigint, Cuenta_origen bigint REFERENCES Cuenta(Nmro_cuenta), Cuenta_destino bigint, Monto int, Fecha timestamp, PRIMARY KEY(Id))"
    create[4] = "CREATE TABLE IF NOT EXISTS Persona_nocliente(Rut varchar(12), Nombre string, Nmro_cuenta bigint, Tipo integer REFERENCES Tipo_cuentas(Id), Banco int REFERENCES Banco(Id), PRIMARY KEY(Rut))"
    create[5] = "CREATE TABLE IF NOT EXISTS Banco(Id int, Nombre string NOT NULL, Ciudad string NOT NULL, PRIMARY KEY (Id))"

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
	var tmp string
	err := db.QueryRow("SELECT pass FROM cliente WHERE rut=$1 AND pass=$2", rut, pass).Scan(&tmp)
	defer disconnect_db()
	switch {
		case err == sql.ErrNoRows:
			return false
		case err != nil:
			return false
		default:
			return true
	}
}

func Account(rut string) (Account_s, error) {
	connect_db()
	var tmp Account_s
	log.Printf("rut -> ", rut)
	fmt.Println(rut)
	row := db.QueryRow("SELECT id, rut_cliente, tipo, saldo FROM cuenta WHERE cuenta.rut_cliente = $1", rut).Scan(&tmp.Id,&tmp.Rut_cliente,&tmp.Tipo,&tmp.Saldo)
	log.Print(row)
	log.Printf("Dentro de acc -> ", tmp.Id)
	disconnect_db()
	//tmp.s = true
	switch {
		case row == sql.ErrNoRows:
	//	tmp.s = false
			return tmp, row
		case row != nil:
	//da		tmp.s = false
			return tmp, row
		default:
			return tmp, row
	}
}
