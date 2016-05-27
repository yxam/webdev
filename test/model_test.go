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
    res :=modelutil.Transferencia("22222222-2","18023904-9",500,1,1)
    if res == false {
        t.Error("Se debio transferir")
    }
}

// no tiene dinero en el saldo para Transferir 
func TestTransferenciaFalse(t *testing.T){
    res :=modelutil.Transferencia("10100100-1","18023904-9",500000,1,1)
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

//deberia mostrar orden por fecha
func TestHistorialTransferenciaTrue1(t *testing.T){
    res :=modelutil.HistorialdeTransferencia("22222222-2","fecha")
    if res == false {
        t.Error("debió mostrar orden por fecha")
    }
}

//deberia mostrar orden por tipo de cuenta
func TestHistorialTransferenciaTrue2(t *testing.T){
    res :=modelutil.HistorialdeTransferencia("22222222-2","tipo_cuenta")
    if res == false {
        t.Error("debió mostrar orden por tipo cuenta")
    }
}

//rut no existe 
func TestHistorialTransferenciaFalse1(t *testing.T){
    res :=modelutil.HistorialdeTransferencia("10100100-4","fecha")
    if res == true {
        t.Error("el rut es inválido no debió haber mostrado")
    }
}

//no existe tipo de orden
func TestHistorialTransferenciaFalse2(t *testing.T){
    res :=modelutil.HistorialdeTransferencia("10100100-1","pipi")
    if res == true {
        t.Error("no existe ese orden, no deberia mostrar nada")
    }
}

//ni el rut ni el tipo de orden existe no deberia mostrar nada
func TestHistorialTransferenciaFalse3(t *testing.T){
    res :=modelutil.HistorialdeTransferencia("10100100-4","pipi")
    if res == true {
        t.Error("rut no existe, ni el tipo de orden")
    }
}

//existe el rut y la cuenta  como originarios
func TestUltimosMovimientosTrue1(t *testing.T){
    res :=modelutil.UltimosMovimientos("22222222-2",2)
    if res == false {
        t.Error("debería mostrar los movimientos al menos del rut originario")
    }
}

//existe el rut y la cuenta como destinatario
func TestUltimosMovimientosTrue2(t *testing.T){
    res :=modelutil.UltimosMovimientos("18023904-9",1)
    if res == false {
        t.Error("deberia mostrar los movimientos al menos del rut destinatario")
    }
}

//no existe rut
func TestUltimosMovimientosFalse1(t *testing.T){
    res :=modelutil.UltimosMovimientos("22222222-4",2)
    if res == true {
        t.Error("no deberia mostrar nada , no existe el rut")
    }
}

//no existe ese tipo de cuenta
func TestUltimosMovimientosFalse2(t *testing.T){
    res :=modelutil.UltimosMovimientos("18023904-9",2)
    if res == true {
        t.Error("no deberia mostrar nada , no existe el tipo de cuenta")
    }
}