package main

import (
	"testing"

	"webdev/cmd/webdev/modelutil"
	//"github.com/franela/goblin" USAREMOS ESTO ;)
)

func TestInit(t *testing.T) {
	flag := modelutil.Init()
	if flag != true {
		t.Error("Mal hecho")
	}
}

//Rut existente
func TestLoginTrue1(t *testing.T) {
    res := modelutil.Login("10100100-1", "0000") 
    if res == false {
    	t.Error("La cuenta si existe")
    }
}

//Rut existente
func TestLoginTrue2(t *testing.T) {
    res := modelutil.Login("18540928-7", "1234") 
    if res == false {
        t.Error("La cuenta si existe")
    }
}

//Rut existente
func TestLoginTrue3(t *testing.T) {
    res := modelutil.Login("18104728-3", "0002") 
    if res == false {
        t.Error("La cuenta si existe")
    }
}

//Rut existente
func TestLoginTrue4(t *testing.T) {
    res := modelutil.Login("0000-0", "0000") 
    if res == false {
        t.Error("La cuenta si existe")
    }
}

//Rut inexistente
func TestLoginFalse1(t *testing.T) {
    res := modelutil.Login("10100100-1", "0001") 
    if res == true {
    	t.Error("La cuenta no existe")
    }
}

//Rut inexistente
func TestLoginFalse2(t *testing.T) {
    res := modelutil.Login("10100900-1", "1111") 
    if res == true {
        t.Error("La cuenta no existe")
    }
}

//Rut inexistente
func TestLoginFalse3(t *testing.T) {
    res := modelutil.Login("10101900-1", "0000") 
    if res == true {
        t.Error("La cuenta no existe")
    }
}

//Rut inexistente
func TestLoginFalse4(t *testing.T) {
    res := modelutil.Login("10100300-1", "0000") 
    if res == true {
        t.Error("La cuenta no existe")
    }
}

//Se realizó la Transferencia
func TestTransferenciaTrue(t *testing.T){
    res :=modelutil.Transferencia("10100100-1","18023904-9", 1000)
    if res == false {
        t.Error("Se debio transferir")
    }
}

// no tiene dinero en el saldo para Transferir 
func TestTransferenciaFalse(t *testing.T){
    res :=modelutil.Transferencia("10100100-1","18023904-9",500000)
    if res == true {
        t.Error ("transfirio lo que no debio ser transferido")
    }
}


//crear cliente exito! tira FAIL pq ya lo creó, pero en el inicio pasó la prueba
func TestCrearClienteTrue(t *testing.T){
    res :=modelutil.IngCliente("22222222-2", 123, 123)
    if res == false {
        t.Error ("debió ser creada")
    }
}

//el rut existe por lo tanto no deberia crear cliente
func TestCrearClienteFalse1(t *testing.T){
    res :=modelutil.IngCliente("18023904-9", 123, 123)
    if res == true {
        t.Error("no debio de ser creada, el rut ya existe")
    }
}

//el rut existe, pero las constraseñas no concuerdan
func TestCrearClienteFalse2(t *testing.T){
    res :=modelutil.IngCliente("33333333-3", 123, 321)
    if res == true {
        t.Error("no debio ser creado, password dont match")
    }
}
//rut existe, password no concuerda, no deberia ser creado por ningun motivo
func TestCrearClienteFalse3(t *testing.T){
    res :=modelutil.IngCliente("18023904-9", 123, 321)
    if res == true{
        t.Error("el rut ya existe, password dont match")
    }
}
