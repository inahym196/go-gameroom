package routing

import (
	controller "go-gameroom/adapter/controller/gin"
	gateway "go-gameroom/adapter/gateway/inmemory"
	"go-gameroom/usecase/interactor"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type GinRouter struct {
	Gin    *gin.Engine
	Melody *melody.Melody
}

func NewGinRouter() Router {
	router := &GinRouter{Gin: gin.Default(), Melody: melody.New()}
	router.setRouting()
	return router
}

func (r *GinRouter) setRouting() {
	roomController := controller.RoomController{
		RepositoryFactory: gateway.NewRoomRepository,
		PortFactory:       interactor.NewRoomLoosePort,
	}
	r.Gin.Static("/public", "./driver/public")
	r.Gin.LoadHTMLGlob("./driver/templates/*")
	r.Gin.GET("/rooms/", func(c *gin.Context) { roomController.GetRooms(c) })
	r.Gin.GET("/rooms/:roomId", func(c *gin.Context) { roomController.GetRoom(c) })
	r.Gin.DELETE("/rooms/:roomId", func(c *gin.Context) { roomController.InitRoom(c) })
}

func (r *GinRouter) Run(addr string) {
	r.Gin.Run(addr)
}
