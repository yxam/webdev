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
func TestLoginTrue(t *testing.T) {
    res := modelutil.Login("10100100-1", "0000") 
    if res == false {
    	t.Error("La cuenta si existe")
    }
}

//Rut inexistente
func TestLoginFalse(t *testing.T) {
    res := modelutil.Login("10100900-1", "0000") 
    if res == true {
    	t.Error("La cuenta no existe")
    }
}