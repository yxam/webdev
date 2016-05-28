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
func ping() {
	err := db.Ping()
	if err != nil {
		log.Println("se calló la conexión")
	}
}

func Init() bool {
	connect_db()

    var create [6]string
	create[0] = "CREATE TABLE IF NOT EXISTS Cliente(Rut varchar(12), Password varchar(4) NOT NULL ,Nombre varchar (255) NOT NULL, Direccion varchar(255) NOT NULL, Comuna varchar(255) NOT NULL, Ciudad varchar (255) NOT NULL, Telefono varchar (20)  NOT NULL, mail varchar (255),PRIMARY KEY(Rut) )"
	create[1] = "CREATE TABLE IF NOT EXISTS Banco (Id int, Nombre varchar(40), Ciudad varchar(255), PRIMARY KEY (Id))"
	create[2] = "CREATE TABLE IF NOT EXISTS Cuenta(Nmro_cuenta bigint, rut_cliente varchar(12) REFERENCES Cliente(Rut), Tipo integer REFERENCES Tipo_cuentas(Id) NOT NULL, Saldo integer NOT NULL, PRIMARY KEY(Nmro_cuenta))"
    create[3] = "CREATE TABLE IF NOT EXISTS Transferencia(Id SERIAL, Cuenta_origen bigint REFERENCES Cuenta(Nmro_cuenta), Cuenta_destino bigint  NOT NULL, Monto int NOT NULL, Fecha timestamp NOT NULL, PRIMARY KEY(Id))"
    create[4] = "CREATE TABLE IF NOT EXISTS Persona_nocliente(Rut varchar(12), Nombre varchar(255), Nmro_cuenta bigint, Tipo integer REFERENCES Tipo_cuentas(Id), Banco int REFERENCES Banco(Id), PRIMARY KEY(Rut))"
    create[5] = "CREATE TABLE IF NOT EXISTS Tipo_cuentas(Id int, Nombre varchar (80) NOT NULL, PRIMARY KEY (Id))"

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
	row := db.QueryRow("SELECT password FROM cliente WHERE rut=$1 AND password=$2", rut, pass)
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

func IngCliente(rut string, nombre string, direccion string, comuna string, ciudad string, telefono int, mail string, pass string, pass2 string) bool{

	connect_db()
	row :=db.QueryRow("SELECT rut FROM cliente WHERE rut = $1", rut)
	//si existe el rut, o la contraseña 1 es distinta  la de verificacion entonces no se ingresa
	if (row.Scan() != sql.ErrNoRows || pass != pass2){
		disconnect_db()
		return false
	}
	//caso de que no exista el rut, y la contraseña 1 es igual a la de verificación se crea cuenta
	_=db.QueryRow("INSERT INTO cliente values ($1,$2,$3,$4,$5,$6,$7,$8)", rut ,pass, nombre, direccion,comuna,ciudad,telefono,mail)
 	disconnect_db()
 	return true
    
}



func Transferencia(rut string, cuenta_o int, cuenta_d int, cantidad int, comentario string) bool{
	var saldo_o, saldo_d int
	var nuevo_o, nuevo_d int
	connect_db()
	//obtiene el saldo de origen
	_=db.QueryRow("select saldo from cuenta where nmro_cuenta=$1",cuenta_o).Scan(&saldo_o)

if (saldo_o > cantidad){
	ping()
	//verifica si la cuenta destino está dentro de los clientes
	 row :=db.QueryRow("SELECT nmro_cuenta from cuenta where nmro_cuenta=$1", cuenta_d)
	if (row.Scan() != sql.ErrNoRows){
		ping()
		//si esta en los clientes le descuenta el monto y updatea su salgo
		_=db.QueryRow("SELECT saldo FROM cuenta WHERE nmro_cuenta=$1",cuenta_d).Scan(&saldo_d)
		nuevo_d = saldo_d + cantidad;
		_=db.QueryRow("UPDATE cuenta SET saldo=$1 where nmro_cuenta= $2",nuevo_d, cuenta_d)
		}	
		ping()
		nuevo_o = saldo_o - cantidad;
		//actualizar saldo del que transfirió
		_=db.QueryRow("UPDATE Cuenta SET saldo=$1 where nmro_cuenta= $2",nuevo_o,cuenta_o)
		//ahora que estan updateados los datos se hace la transferencia
		fecha:=time.Now()
		_=db.QueryRow("INSERT INTO Transferencia VALUES ($1,$2,$3,$4,$5)",cuenta_o, cuenta_d, cantidad, fecha, comentario)
		//si no hay error y la transferencia fue un exito!
		disconnect_db()
		return true
	}

	//en caso que no se pueda transferir, retorna falso cuando no tiene saldo suficiente para transferir
disconnect_db()
	return false
}


func HistorialdeTransferencia(cuenta int) bool{

	connect_db()
	//va a comprobar si la cuenta esta en las transferencias
	row :=db.QueryRow("SELECT * FROM Transferencia WHERE cuenta_origen = $1", cuenta)


	//si no existe  es porque no ha hecho ninguna transferencia
		if (row.Scan() == sql.ErrNoRows){
		disconnect_db()
		return false
		}
 
	_=db.QueryRow("SELECT * from transferencia WHERE cuenta_origen=$1 ORDER BY fecha",cuenta)
	disconnect_db()
	return true
}


func UltimosMovimientos(cuenta int) bool {

	connect_db()
	//comprueba que la cuenta haya tenido transacciones 
	row :=db.QueryRow("SELECT * FROM transferencia WHERE cuenta_origen =$1 OR cuenta_destino =$1", cuenta)

	//si no existe ninguna de esas variables que retorne falso
	if (row.Scan() == sql.ErrNoRows){
		disconnect_db()
		return false
	}
	//caso que tenga transacciones
	_=db.QueryRow("SELECT * from transferencia WHERE cuenta_origen =$1 OR cuenta_destino =$1 ORDER BY fecha",cuenta)
	disconnect_db()
	return true
}
