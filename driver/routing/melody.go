package routing

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func GetMRouter() *melody.Melody {
	m := melody.New()
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		println(string(msg))
	})
	return m
}

func UpgradeHandleFunc(m *melody.Melody) func(*gin.Context) {
	return func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	}
}
