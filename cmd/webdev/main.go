package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	//ConfigRuntime()
	//StartWorkers()
	StartGin(port)
}

func StartGin(port string) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), ipNotAllowed())
	log.Print(router)
	router.LoadHTMLGlob("resources/*.html")
	router.Static("/static", "resources/static")
	router.GET("/", index)
	router.GET("/marii", marii)
	router.GET("/createdb", createdb)
	router.POST("/processLogin", processLogin)

	if err := router.Run(":" + port); err != nil {
		log.Printf("error listening on port "+port+": %v", err)
	}
}
