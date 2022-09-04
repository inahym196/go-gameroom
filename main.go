package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	engine := gin.Default()
	m := melody.New()
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
	engine.GET("/rooms/:id/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	engine.Run("127.0.0.1:3000")
}
