package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"trunk_exem/employment_examination/golang/routes"
)

func main() {
	app :=gin.Default()

	routes.Setup(app)

	if err :=app.Run(":5000"); err !=nil{
		log.Fatalln(err.Error())
	}
}

