package routing

import (
	"fmt"
	controller "go-gameroom/adapter/controller/gin"
	"go-gameroom/usecase/port"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type Router interface {
	Run(addr string)
}

type GinRouter struct {
	gin            *gin.Engine
	melody         *melody.Melody
	userController controller.UserController
	gameController controller.GameController
}

func NewGinRouter(
	userController controller.UserController,
	gameController controller.GameController,
) Router {
	router := &GinRouter{
		gin:            gin.Default(),
		melody:         melody.New(),
		userController: userController,
		gameController: gameController,
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
	_games, err := r.gameController.GetGames()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	var games = make(map[string]port.GetGameResponse)
	for i, game := range _games {
		s := strconv.Itoa(i)
		games[s] = *game
	}
	fmt.Printf("%#v", games)
	c.HTML(http.StatusOK, "lobby.tmpl", gin.H{
		"userName": c.GetString("userName"),
		"Games":    games,
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
