package routing

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *GinRouter) CheckLogin(c *gin.Context) {
	c.Set("userName", "Guest")
	sessionId, err := c.Cookie("sessionId")
	if err != nil {
		fmt.Println(err)
		c.Next()
		return
	}
	userName, err := r.UserController.GetUserName(sessionId)
	if err != nil {
		fmt.Println(err)
		c.Next()
		return
	}
	c.Set("userName", userName)
	c.Next()
	return
}

func (r *GinRouter) MustLogout(c *gin.Context) {
	_, err := c.Cookie("sessionId")
	if err != nil {
		fmt.Println(err)
		c.Next()
		return
	}
	c.Redirect(http.StatusFound, "/lobby")
	return
}
