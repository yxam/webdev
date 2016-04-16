package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"fmt"
)

func main (){

//ingresar saldo
Saldo()

//ingresar cliente
//IngCliente()

//login
//Login()

}



//consultar saldo
func Saldo(){
	db, err := sql.Open("postgres", "postgres://tbllgrkjejpwzv:e3D-VEc5BmjTyw6pESuJnzgQAo@ec2-54-221-249-201.compute-1.amazonaws.com:5432/dcvc2lb7meb7j5")
	if err != nil {
		log.Println(err)}
	defer db.Close()

	var numerosaldo int
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