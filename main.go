package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Static("/public", "./public")
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Go GameRoom",
		})
	})
	engine.GET("/rooms/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.HTML(http.StatusOK, "room.tmpl", gin.H{
			"title": "Go GameRoom",
			"id":    id,
		})
	})
	engine.Run("127.0.0.1:3000")
}
