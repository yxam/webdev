package modelutil

import (
	"database/sql"
	"log"
	//"os"
	_ "github.com/lib/pq"
	"time"
)

var db *sql.DB = nil
var err error

type Information struct {
	rut  string 
	pass string //`form:"pass"`// json:"pass" binding:"required"`
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
    create[3] = "CREATE TABLE IF NOT EXISTS Transferencia(rut_origen varchar(12) REFERENCES cliente(rut), cuenta_origen integer, rut_destino varchar(12) NOT NULL, cuenta_destino integer, monto integer NOT NULL,fecha timestamp,PRIMARY KEY (rut_origen,fecha))" 
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
	row := db.QueryRow("SELECT pass FROM cliente WHERE rut=$1 AND pass=$2", rut, pass)
	if row.Scan() == sql.ErrNoRows {
		return false
	} 
	return true
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

func IngCliente(rut_n string, pass int, pass2 int) bool{

	connect_db()
	row :=db.QueryRow("SELECT rut FROM cliente WHERE cliente.rut = $1", rut_n)
	//si existe el rut, o la contraseña 1 es distinta  la de verificacion entonces no se ingresa
	if (row.Scan() != sql.ErrNoRows || pass != pass2){
		disconnect_db()
		return false
	}
	//caso de que no exista el rut, y la contraseña 1 es igual a la de verificación se crea cuenta
	_=db.QueryRow("INSERT INTO cliente values ($1,$2)", rut_n, pass)
 	disconnect_db()
 	return true
    
}



func Transferencia(rut_o string, rut_d string, cantidad int, cuenta_o int, cuenta_d int) bool{
	var saldo_o, saldo_d int
	var nuevo_o, nuevo_d int
	connect_db()
	//obtener saldos de origen y destino
	_=db.QueryRow("select saldo from cuenta where rut_cliente=$1 AND tipo=$2",rut_o, cuenta_o).Scan(&saldo_o)
	_=db.QueryRow("select saldo from cuenta where rut_cliente=$1 AND tipo=$2",rut_d, cuenta_d).Scan(&saldo_d)
	disconnect_db()
    
	//comprobar que se puede efectuar la transferencia
	
	
	if (saldo_o > cantidad){

		nuevo_o = saldo_o - cantidad;
		nuevo_d = saldo_d + cantidad;

	//update a las tablas
		connect_db()
		//actualizar saldo del que transfirió
		_=db.QueryRow("UPDATE Cuenta SET saldo=$1 where rut_cliente= $2 AND tipo=$3",nuevo_o, rut_o, cuenta_o)

		//actualizar saldo del destinatario
		_=db.QueryRow("UPDATE cuenta SET saldo=$1 where rut_cliente= $2 AND tipo=$3",nuevo_d, rut_d, cuenta_d)
		


		//ahora que estan updateados los datos se hace la transferencia
		fecha:=time.Now()
		_=db.QueryRow("INSERT INTO Transferencia VALUES ($1,$2,$3,$4,$5,$6)",rut_o, cuenta_o, rut_d, cuenta_d, cantidad, fecha)
		


		//si no hay error y la transferencia fue un exito!
		disconnect_db()
		return true
	}

	//en caso que no se pueda transferir, retorna falso cuando no tiene saldo suficiente para transferir

	return false
}

/*
func HistorialdeTransferencia(rut_o string, orden string) bool{

	connect_db()
	//va a comprobar si el rut existe en la tabla transferencia
	row :=db.QueryRow("SELECT * FROM Transferencia WHERE rut_origen = $1", rut_o)


	//si no existe el rut en la tabla porque no ha transferido o ingreso mal un orden, no se ingresa
		if (row.Scan() == sql.ErrNoRows || orden != "fecha" || orden != "tipo_cuenta"){
			log.Println(orden)
		disconnect_db()
		return false
		}
 
	_=db.QueryRow("SELECT * from transferencia WHERE rut_origen=$1 ORDER BY $2 ",rut_o, orden)
	disconnect_db()
	return true
}


func UltimosMovimientos(rut string, numero_cuenta int) bool {

	connect_db()
	//comprueba que el rut y el numero de cuenta existan
	row :=db.QueryRow("SELECT * FROM transferencia WHERE (rut_origen=$1 OR rut_destino=$1) AND (cuenta_origen= $2 OR cuenta_destino= $2)", rut, numero_cuenta)

	//si no existe ninguna de esas variables que retorne falso
	if (row.Scan() == sql.ErrNoRows){
		disconnect_db()
		return false
	}

	_=db.QueryRow("SELECT * from transferencia WHERE (rut_origen=$1 OR rut_destino=$1) AND (cuenta_origen=$2 OR cuenta_destino= $2) ORDER BY fecha",rut, numero_cuenta)
	disconnect_db()
	return true
}
*/
