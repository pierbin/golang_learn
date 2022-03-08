package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"golang_learn/pkgusages/gin_gorilla/ws"
)

func main() {
	go ws.Manager.Start()
	r := gin.Default()
	r.GET("/json", ws.JsonApi)
	r.GET("/text", ws.TextApi)
	r.GET("/chat", ws.ChatHandler)
	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// static files
	r.Use(static.Serve("/", static.LocalFile("./", true)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./index.html")
	})

	_ = r.Run(":8000")
}
