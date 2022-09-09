package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {

	url := "http://localhost:8000/api/v1/boards/0/"

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

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		byteArray, _ := io.ReadAll(resp.Body)
		m.Broadcast(byteArray)
	})
	engine.Run("127.0.0.1:3000")
}
