package main

import "github.com/gin-gonic/gin"
import "log"
import "strconv"

//
//Falta el acceso a un archivo que permita verificar en caché
//las IP para USA.


// Función que funciona como middleware prohibiendo entrada a alguna IP de USA.
func ipNotAllowed() gin.HandlerFunc {
	return func(c *gin.Context) {
		tmp := c.ClientIP()[0:3]
		var n int
		//var err error
		if tmp[2] == '.' {
			n, _ = strconv.Atoi(tmp[0:2]) //queda pendiente el error
		}
		//log.Print(n)
		//if err != nil {
		if n <= 127 {
			log.Printf("IP denied")
			c.AbortWithStatus(403)
		}
		//} else {
		//	c.AbortWithStatus(403)
		//}
	}
}
