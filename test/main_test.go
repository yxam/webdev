package main

import (
	"testing"
	"net/http"
	"log"
	"strconv"

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
	if res.StatusCode != http.StatusNoContent {
		t.Error("Error en funcion processLogin")
	}
}

func TestProcessLoginParametrosValidosIncorrectos(t *testing.T) {
	item := Item {
		rut: "123", // Este usuario no existe en la base de datos
		pass: "123",
	}

	res, _ := goreq.Request{
		Method : "POST",
		Uri : "https://abbanks.herokuapp.com/processLogin",
		QueryString: item,
	}.Do()
	log.Print("- StatusCode -> " + strconv.Itoa(res.StatusCode))
	
}