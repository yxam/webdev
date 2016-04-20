package main

import (
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

type information struct {
	rut  string //`form:"rut"`// json:"rut" binding:"required"
	pass string //`form:"pass"`// json:"pass" binding:"required"`
}

type response struct {
	rut string
	pass string
	saldo int
	message string
}

func processLogin(c *gin.Context) {
	var inf_tmp modelutil.Information
	inf_tmp.Rut = c.PostForm("rut")
	inf_tmp.Pass = c.PostForm("pass")
	
	if inf_tmp.Rut != "" && inf_tmp.Pass != "" {
		state := modelutil.Login(inf_tmp.Rut, inf_tmp.Pass)
		if state {
			account := modelutil.Account(inf_tmp)
			if account != nil {
				c.JSON(http.StatusOK, account)
			} else {
				c.JSON(http.StatusForbidden, gin.H{})
			}
		} else {
				c.HTML(http.StatusBadRequest, "index.html", gin.H{"message":"User and/or pass is invalid"})
		}
	} else {
		c.AbortWithStatus(http.StatusNoContent) //Debe salirCREO
	}
}

func createdb(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://tbllgrkjejpwzv:e3D-VEc5BmjTyw6pESuJnzgQAo@ec2-54-221-249-201.compute-1.amazonaws.com:5432/dcvc2lb7meb7j5")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message":"Error en la db"})
	}
	defer db.Close()

    //var create []string
	//
	create, err = db.Prepare("CREATE TABLE IF NOT EXISTS Cliente (rut varchar(12), pass varchar(4) NOT NULL, PRIMARY KEY(rut)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message":"Error creando tabla"})
	}
	//create[1] = "CREATE TABLE IF NOT EXISTS Banco (id serial, nombre varchar(50) NOT NULL, PRIMARY KEY (id))"
	////serial = (int) auto_increment
    //create[2] = "CREATE TABLE IF NOT EXISTS Cuenta(id bigint, rut_cliente varchar(12) REFERENCES cliente(rut), tipo integer NOT NULL, saldo integer NOT NULL"
    ////id = numero de cuenta, por eso bigint y no serial que es auto incremental. 
    //create[3] ="CREATE TABLE IF NOT EXISTS Transferencia(rut_origen varchar(12) REFERENCES cliente(rut), rut_destino varchar(12) NOT NULL,monto integer NOT NULL, fecha timestamp,PRIMARY KEY (rut_origen,fecha))" 
    //timestamp, guarda fecha y hora
//    var length = cap(create)
//    i := 0
//    for i < length { 
//	    _, err := db.Exec(create[i])    
//	    if err != nil {
//			//disconnect_db()
//	        db.Close()
//	        return false
//	    }
//	    i++
//	}
	res, err := create.Exec()
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"Message":"Error ejecutando consulta"})
	}
	//db.Close()
	//return true
//
//	//flag := modelutil.Init()
//	//if flag {
//	//	c.JSON(http.StatusOK, gin.H{"message":"database created!"})
//	//} else {
//	//	c.JSON(http.StatusInternalServerError, gin.H{"message":"database was created previously"})
	//}
}


func marii(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://tbllgrkjejpwzv:e3D-VEc5BmjTyw6pESuJnzgQAo@ec2-54-221-249-201.compute-1.amazonaws.com:5432/dcvc2lb7meb7j5")
	if err != nil {
		c.JSON(400, gin.H{"Chupalo":"maxi"})
	}
	defer db.Close()

	var numerosaldo float64
	var rut string
	
	rut = "123"

	err = db.QueryRow("select saldo from cuenta where cuenta.rut_cliente=$1",rut).Scan(&numerosaldo)
	if err != nil && err !=sql.ErrNoRows {
		c.JSON(400, gin.H{"Message":"Aprovecha AWEONAO"})
	}
	c.JSON(200, gin.H{"Message":numerosaldo})

}


