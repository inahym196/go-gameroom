package controller

import (
	model_redis "go-gameroom/model/redis"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func getLobby(c *gin.Context) {
	cookieKey := "LoginUserIdKey"
	userName := model_redis.GetSession(c, cookieKey)
	if userName == nil {
		userName = "Guest"
	}
	c.HTML(http.StatusOK, "lobby.tmpl", gin.H{
		"title":    "Go GameRoom",
		"userName": userName,
	})
}

func getRoom(c *gin.Context) {
	id := c.Param("id")
	userName, exists := c.Get("authedUser")
	if exists == false {
		userName = "Guest"
	}
	c.HTML(http.StatusOK, "room.tmpl", gin.H{
		"title":    "Go GameRoom",
		"userName": userName,
		"id":       id,
	})
}

func getWS(m *melody.Melody) func(c *gin.Context) {
	return func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	}
}

func checkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieKey := "LoginUserIdKey"
		userName := model_redis.GetSession(c, cookieKey)
		if userName == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		} else {
			c.Set("authedUser", userName)
			c.Next()
		}
	}
}

func checkLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieKey := "LoginUserIdKey"
		userName := model_redis.GetSession(c, cookieKey)
		if userName != nil {
			c.Redirect(http.StatusFound, "/lobby")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func GetRouter(m *melody.Melody) *gin.Engine {
	router := gin.Default()
	router.Static("/public", "./public")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", getLobby)
	router.GET("/lobby", getLobby)
	loginCheckGroup := router.Group("/", checkLogin())
	{
		loginCheckGroup.GET("/rooms/:id", getRoom)
		loginCheckGroup.GET("/rooms/:id/ws", getWS(m))
		loginCheckGroup.GET("/logout", getLogout)
	}
	logoutCheckGroup := router.Group("/", checkLogout())
	{
		logoutCheckGroup.GET("/login", getLogin)
		logoutCheckGroup.POST("/login", postLogin)
	}
	return router
}
