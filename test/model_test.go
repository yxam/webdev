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