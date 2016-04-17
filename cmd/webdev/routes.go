package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func rateLimit(c *gin.Context) {

	ip := c.ClientIP()
	value := int(ips.Add(ip, 1))
	if value%50 == 0 {
		fmt.Printf("ip: %s, count: %d\n", ip, value)
	}
	if value >= 200 {
		if value%200 == 0 {
			fmt.Println("ip blocked")
		}
		c.Abort()
		c.String(503, "you were automatically banned :)")
	}
}



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
	var inf_tmp information
	inf_tmp.rut = c.PostForm("rut")
	inf_tmp.pass = c.PostForm("pass")
	inf.rut = "10100100-1"
	inf.pass = "1234"
	
	if inf_tmp.rut != "" && inf_tmp.pass != "" {
		state = login(inf_tmp)
		if state {
			account := account(inf_tmp)
			if account != nil {
				c.JSON(http.StatusOK, account)
			} else {
				c.JSON(http.StatusForbidden, gin.H{})
			}
		} else {
				c.HTML(http.StatusBadRequest, "index.html", gin.H{"User and/or pass is invalid"})
		}
	} else {
		c.AbortWithStatus(http.StatusNoContent) //Debe salirCREO
	}
}


