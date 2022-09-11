package controller

import (
	model_redis "go-gameroom/model/redis"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func getLobby(c *gin.Context) {
	userName := c.GetString("authedUser")
	c.HTML(http.StatusOK, "lobby.tmpl", gin.H{
		"userName": userName,
	})
}

func getRoom(c *gin.Context) {
	id := c.Param("id")
	userName := c.GetString("authedUser")
	c.HTML(http.StatusOK, "room.tmpl", gin.H{
		"userName": userName,
		"id":       id,
	})
}

func getWS(m *melody.Melody) func(c *gin.Context) {
	return func(c *gin.Context) {
		userName := c.GetString("authedUser")
		m.HandleRequestWithKeys(c.Writer, c.Request, map[string]interface{}{"userName": userName})
	}
}

func checkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieKey := "LoginUserIdKey"
		userName := model_redis.GetSession(c, cookieKey)
		if userName == nil {
			userName = "Guest"
		}
		c.Set("authedUser", userName)
		c.Next()
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

	loginCheckGroup := router.Group("/", checkLogin())
	{
		loginCheckGroup.GET("/", getLobby)
		loginCheckGroup.GET("/lobby", getLobby)
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
