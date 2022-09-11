package controller

import (
	model_redis "go-gameroom/model/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func postLogin(c *gin.Context) {
	userName := c.PostForm("userName")

	if userName != "first-user" && userName != "draw-user" {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	cookieKey := "LoginUserIdKey"
	model_redis.NewSession(c, cookieKey, userName)
	c.Redirect(http.StatusFound, "/lobby")
}

func getLogout(c *gin.Context) {
	cookieKey := "LoginUserIdKey"
	model_redis.DeleteSession(c, cookieKey)
	c.Redirect(http.StatusFound, "/lobby")
}
