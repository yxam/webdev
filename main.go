package main

import (
	"runtime"
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
                log.Fatal("$PORT must be set")
        }

	ConfigRuntime()
	StartWorkers()
	StartGin(port)
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	//fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartWorkers() {
	go statsWorker()
}

func StartGin(port string){
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.LoadHTMLGlob("../src/main/resources/*.html")
	router.Static("/static", "../src/main/resources/static")
	router.GET("/index", index)
	router.POST("/processInfo", processInfo)
	
	router.Run(":"+port)
}
