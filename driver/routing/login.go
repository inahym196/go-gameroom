package routing

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *GinRouter) MustLogin(c *gin.Context) {
	sessionId, err := c.Cookie("sessionId")
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	fmt.Println("MustLogin: ", sessionId)
	userName, err := r.UserController.GetUserName(sessionId)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	c.Set("userName", userName)
	c.Next()
	return
}

func (r *GinRouter) MustLogout(c *gin.Context) {
	sessionId, err := c.Cookie("sessionId")
	_, err = r.UserController.GetUserName(sessionId)
	if err != nil {
		fmt.Println(err)
		c.Next()
		return
	}
	c.Redirect(http.StatusFound, "/lobby")
	return
}

func (r *GinRouter) PostLogin(c *gin.Context) {
	userName := c.PostForm("userName")
	sessionId, err := r.UserController.CreateSession(userName)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	c.SetCookie("sessionId", sessionId, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/lobby")
}

func (r *GinRouter) GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

func (r *GinRouter) GetLogout(c *gin.Context) {
	sessionId, _ := c.Cookie("sessionId")
	r.UserController.DeleteSession(sessionId)
	c.SetCookie(sessionId, "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}
