package main

import (
	"testing"

	"webdev/cmd/webdev/modelutil"
	//"github.com/franela/goblin" USAREMOS ESTO ;)
)

func TestInit(t *testing.T) {
	flag := modelutil.Init()
	if flag == true {
		t.Error("Mal hecho")
	}
}

//Max
func TestLogin(t *testing.T) {
    var rut string
    var pass string
    rut = "10100100-1"
    pass = "0000"
    res := modelutil.Login(rut, pass)
    if res == false {
    	t.Error("La cuenta si existe")
    }
}