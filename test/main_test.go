package main

import (
	"testing"

	"github.com/franela/goreq"
	//"github.com/franela/goblin" USAREMOS ESTO ;)
)

func TestPagina(t *testing.T) {
	_, err := goreq.Request{
		Uri : "https://abbanks.herokuapp.com/",
	}.Do()
	if err != nil {
		t.Error("Error en abrir index.html")
	}
}