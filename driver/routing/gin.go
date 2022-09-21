package routing

import (
	controller "go-gameroom/adapter/controller/gin"
	"go-gameroom/usecase/port"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type GinRouter struct {
	Gin            *gin.Engine
	Melody         *melody.Melody
	RoomController controller.RoomController
	UserController controller.UserController
}

func NewGinRouter(roomController controller.RoomController, userController controller.UserController) Router {
	router := &GinRouter{
		Gin:            gin.Default(),
		Melody:         melody.New(),
		RoomController: roomController,
		UserController: userController,
	}
	router.setRouting()
	return router
}

func (r *GinRouter) setRouting() {
	r.Gin.Static("/public", "./driver/public")
	r.Gin.LoadHTMLGlob("./driver/templates/*")
	loginCheckGroup := r.Gin.Group("/", r.MustLogin)
	{
		loginCheckGroup.GET("/", r.GetLobby)
		loginCheckGroup.GET("/lobby", r.GetLobby)
		loginCheckGroup.GET("/rooms/:roomId", r.GetRoom)
		loginCheckGroup.GET("/logout", r.GetLogout)
	}
	logoutCheckGroup := r.Gin.Group("/", r.MustLogout)
	{
		logoutCheckGroup.GET("/login", r.GetLogin)
		logoutCheckGroup.POST("/login", r.PostLogin)
	}
}

func (r *GinRouter) Run(addr string) {
	r.Gin.Run(addr)
}

func (r *GinRouter) GetLobby(c *gin.Context) {
	_rooms, err := r.RoomController.GetRooms()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	var rooms = make(map[string]port.RoomDto)
	for i, room := range _rooms {
		rooms[i] = *room
	}
	c.HTML(http.StatusOK, "lobby.tmpl", gin.H{
		"userName": c.GetString("userName"),
		"Rooms":    rooms,
	})

}
func (r *GinRouter) GetRoom(c *gin.Context) {
	roomIdStr := c.Param("roomId")
	roomId, _ := strconv.Atoi(roomIdStr)
	c.HTML(http.StatusOK, "room.tmpl", gin.H{
		"userName": c.GetString("userName"),
		"id":       roomId,
	})

}
