package main

import (
	"net/http"
	"webdev/cmd/webdev/modelutil"

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
	var inf information
	var inf_tmp modelutil.Information
	inf_tmp.Rut = c.PostForm("rut")
	inf_tmp.Pass = c.PostForm("pass")
	
	if inf_tmp.Rut != "" && inf_tmp.Pass != "" {
		state := modelutil.Login(inf_tmp)
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
	if modelutil.Init() {
		c.JSON(http.StatusOK, gin.H{"message":"database created!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message":"database was created previously"})
	}
}