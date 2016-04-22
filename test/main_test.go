package main

import (
	"testing"
	"net/http"
	"fmt"	

	"github.com/franela/goreq"
	//"github.com/franela/goblin" USAREMOS ESTO ;)
)

type Item struct {
        rut string
        pass string
}

func TestIndex(t *testing.T) {
	_, err := goreq.Request{
		Uri : "https://abbanks.herokuapp.com/",
	}.Do()
	if err != nil {
		t.Error("Error en abrir index.html")
	}
}

func TestProcessLoginParametrosVacios(t *testing.T) {
	item := Item {
	        rut: "",
	        pass: "",
	}

	res, _ := goreq.Request{
		Method : "POST",
		Uri : "https://abbanks.herokuapp.com/processLogin",
		QueryString: item,
	}.Do()
	if res.StatusCode != http.StatusMovedPermanently {
		t.Error("Error en funcion processLogin")
	}
}

func TestProcessLoginParametrosValidosIncorrectos(t *testing.T) {
	item := Item {
		rut: "10100100-1", // Este usuario no existe en la base de datos
		pass: "0000",
	}

	res, _ := goreq.Request{
		Method : "POST",
		Uri : "https://abbanks.herokuapp.com/processLogin",
		QueryString: item,
	}.Do()
	fmt.Println("Res -> ", res.StatusCode)
	var m map[string]string
	res.Body.FromJsonTo(&m)
	fmt.Println(m)
	if res.StatusCode != http.StatusInternalServerError {
		t.Error("Pagina debio no pasar")
	}
}