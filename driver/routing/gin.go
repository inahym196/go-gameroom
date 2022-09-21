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
	gin            *gin.Engine
	melody         *melody.Melody
	roomController controller.RoomController
	userController controller.UserController
}

func NewGinRouter(roomController controller.RoomController, userController controller.UserController) Router {
	router := &GinRouter{
		gin:            gin.Default(),
		melody:         melody.New(),
		roomController: roomController,
		userController: userController,
	}
	router.setRouting()
	router.setMelodyRouting()
	return router
}

func (r *GinRouter) setRouting() {
	r.gin.Static("/public", "./driver/public")
	r.gin.LoadHTMLGlob("./driver/templates/*")
	loginCheckGroup := r.gin.Group("/", r.MustLogin)
	{
		loginCheckGroup.GET("/", r.GetLobby)
		loginCheckGroup.GET("/lobby", r.GetLobby)
		loginCheckGroup.GET("/rooms/:roomId", r.GetRoom)
		loginCheckGroup.GET("/rooms/:roomId/ws", r.UpgradeHandleFunc)
		loginCheckGroup.GET("/logout", r.GetLogout)
	}
	logoutCheckGroup := r.gin.Group("/", r.MustLogout)
	{
		logoutCheckGroup.GET("/login", r.GetLogin)
		logoutCheckGroup.POST("/login", r.PostLogin)
	}
}

func (r *GinRouter) Run(addr string) {
	r.gin.Run(addr)
}

func (r *GinRouter) GetLobby(c *gin.Context) {
	_rooms, err := r.roomController.GetRooms()
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
