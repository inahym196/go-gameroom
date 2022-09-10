package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func GetLobby(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Go GameRoom",
	})
}
func GetRoom(c *gin.Context) {
	id := c.Param("id")
	c.HTML(http.StatusOK, "room.tmpl", gin.H{
		"title": "Go GameRoom",
		"id":    id,
	})
}
func GetWS(m *melody.Melody) func(c *gin.Context) {
	return func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	}
}

func GetRouter(m *melody.Melody) *gin.Engine {
	router := gin.Default()
	router.Static("/public", "./public")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", GetLobby)
	router.GET("/rooms/:id", GetRoom)
	router.GET("/rooms/:id/ws", GetWS(m))
	return router
}
