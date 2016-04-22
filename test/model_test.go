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
    res := modelutil.Login("10100100-1", "0000")
    if res == false {
    	t.Error("La cuenta si existe")
    }
}