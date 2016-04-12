package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	//fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartWorkers() {
	go statsWorker()
}

func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.LoadHTMLGlob("../src/main/resources/*.html")
	router.Static("/static", "../src/main/resources/static")
	router.GET("/index", index)
	router.POST("/processInfo", processInfo)
	
	router.Run(":8000")
}
