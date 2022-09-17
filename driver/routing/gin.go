package routing

import (
	"github.com/gin-gonic/gin"
)

func GetRouter(upgradeHandleFunc func(*gin.Context)) *gin.Engine {
	router := gin.Default()
	router.Static("/public", "./driver/public")
	router.LoadHTMLGlob("./driver/templates/*")
	return router
}
