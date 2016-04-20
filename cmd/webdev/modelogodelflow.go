package consultas

/* Maria Cristina Binfa*/

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"fmt"
	//"time"
)

func main (){

//ingresar saldo
//Saldo()

//ingresar cliente
//IngCliente()

//login
//Login()

//hacertransferencia
HacerTransferencia()
}



//consultar saldo

func Saldo(){
	db, err := sql.Open("postgres", "postgres://tbllgrkjejpwzv:e3D-VEc5BmjTyw6pESuJnzgQAo@ec2-54-221-249-201.compute-1.amazonaws.com:5432/dcvc2lb7meb7j5")
	if err != nil {
		log.Println(err)}
	defer db.Close()

	var numerosaldo float64
	var rut int
	fmt.Println("ingrese rut")
	fmt.Scan(&rut)


	err = db.QueryRow("select saldo from cuenta where cuenta.rut_cliente=$1",rut).Scan(&numerosaldo)
if err != nil && err !=sql.ErrNoRows {
	log.Println(err)
}
log.Println("su saldo es:",numerosaldo)}



//ingresar cliente
func IngCliente(){
	db, err := sql.Open("postgres", "postgres://tbllgrkjejpwzv:e3D-VEc5BmjTyw6pESuJnzgQAo@ec2-54-221-249-201.compute-1.amazonaws.com:5432/dcvc2lb7meb7j5")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	var r int
	var p int
	var m string
	var n string
	var a string

	fmt.Println("ingrese rut,password,mail,nombre,apellido")
	_, err = fmt.Scan(&r, &p, &m, &n,&a)
	if err != nil {
		log.Print("Error scanf")
		return
	}
	

	stmt, err := db.Prepare("INSERT INTO cliente values ($1,$2,$3,$4,$5)")
	if err != nil{
	log.Fatal(err)}
	defer stmt.Close()
	res, err := stmt.Exec(r,p,m,n,a)
	if err != nil{
		log.Fatal(err)}

	id, err := res.LastInsertId()
	fmt.Println(id)}




//LOGIN
func Login(){
	var contraseña int
	var pass int
	var id int
	fmt.Println("ingrese rut")
	fmt.Scan(&id)
	fmt.Println("ingrese constraseña")
	fmt.Scan(&contraseña)

	db, err := sql.Open("postgres", "postgres://tbllgrkjejpwzv:e3D-VEc5BmjTyw6pESuJnzgQAo@ec2-54-221-249-201.compute-1.amazonaws.com:5432/dcvc2lb7meb7j5")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()



	err = db.QueryRow("select password from cliente where cliente.rut=$1",id).Scan(&pass)
if err != nil && err !=sql.ErrNoRows {
	log.Println(err)
}
	if pass== contraseña {
		fmt.Println("EEXITO")
	}	



}


// Transferencias

func HacerTransferencia(){
	db, err := sql.Open("postgres", "postgres://tbllgrkjejpwzv:e3D-VEc5BmjTyw6pESuJnzgQAo@ec2-54-221-249-201.compute-1.amazonaws.com:5432/dcvc2lb7meb7j5")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	
var rut_o int 
var rut_d int
var saldo_o, saldo_d float64
var nuevo_o, nuevo_d float64 //nuevo saldo
var cant float64 

fmt.Println("ingrese su rut")
fmt.Scan(&rut_o)
fmt.Println("ingrese rut destinatario")
fmt.Scan(&rut_d)
fmt.Println("ingrese monto")
fmt.Scan(&cant)

err = db.QueryRow("select saldo from cuenta where rut_cliente=$1",rut_o).Scan(&saldo_o)
if err != nil && err !=sql.ErrNoRows {
log.Println(err)}

err = db.QueryRow("select saldo from cuenta where rut_cliente=$2",rut_d).Scan(&saldo_d)//saco el saldo de la cuenta de destinatario
if err != nil{
log.Fatal(err)}

fmt.Println("saldo original")
fmt.Println(saldo_o)
fmt.Println(saldo_d)

//caso que tenga saldo para cubrir el monto a transferir

if (saldo_o>cant){

//operaciones matematicas complejas

nuevo_o = saldo_o - cant;
nuevo_d = saldo_d + cant;

fmt.Println(nuevo_o)
fmt.Println(nuevo_d)


// ahora hago un update de la información de los clientes

//actualiza el saldo de origen
/*stmt, err = db.Prepare("update cuenta set saldo=? where rut_cliente=?")
if err != nil {
		log.Println(err)}
res, err = stmt.Exec(nuevo_o, rut_o) 
if err != nil{
	log.Fatal(err)}
affect, err := res.RowsAffected()
fmt.Println(affect)

//actualiza el saldo de destino

stmt, err = db.Prepare("update cuenta set saldo=? where rut_cliente=?")
if err != nil {
		log.Println(err)}
res, err = stmt.Exec(nuevo_d, rut_d) 
if err != nil{
	log.Fatal(err)}
defer stmt.Close()
affect, err := res.RowsAffected()
fmt.Println(affect)*/
//ahora que esta todo actualizado mandamos el insert a la tabla transferencias

/*stmt, err := db.Prepare("INSERT into transferencias values (?,?,?,?)")
	if err != nil {
		log.Println(err)}
    

    res, err := stmt.Exec(rut_o, rut_d, cant,time.Now().Format(time.RFC850))
    if err != nil{
	log.Fatal(err)}
	defer stmt.Close()

    id, err := res.LastInsertId()
    if err != nil {
		log.Println(err)}

    fmt.Println(id)*/
}

//en caso que no tenga la cantidad necesaria
//if ( saldo_o<cant){
//	println(" tay pato")}

}









/*func Vertransferencias(){
	db, err := sql.Open("postgres", "postgres://tbllgrkjejpwzv:e3D-VEc5BmjTyw6pESuJnzgQAo@ec2-54-221-249-201.compute-1.amazonaws.com:5432/dcvc2lb7meb7j5")
	if err != nil {
		log.Println(err)}
	defer db.Close()

	var rut int
	fmt.Println("ingrese rut")
	fmt.Scan(&rut)


	err = db.QueryRow("select * from transferencias where transferencias.origen=$1",rut)
if err != nil && err !=sql.ErrNoRows {
	log.Println(err)}



}*/