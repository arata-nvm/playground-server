package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/visket-lang/playground/handler"
	"io"
	"log"
	"os"
)

func main() {
	initLog()

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	r.Use(cors.New(config))

	r.POST("/", handler.PostCode)
	r.Run()
}

func initLog() {
	logFile, err := os.Create("./log/request.log")
	if err != nil {
		log.Fatal(err)
	}

	errLogFile, err := os.Create("./log/error.log")
	if err != nil {
		log.Fatal(err)
	}

	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, errLogFile)
	log.SetOutput(gin.DefaultWriter)
}
