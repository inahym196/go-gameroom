package routing

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func (r *GinRouter) UpgradeHandleFunc(c *gin.Context) {
	r.melody.HandleRequest(c.Writer, c.Request)
}

func (r *GinRouter) setMelodyRouting() {
	r.melody.HandleMessage(func(s *melody.Session, msg []byte) {
		println(string(msg))
	})
}
