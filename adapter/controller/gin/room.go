package controller

import (
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/port"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	RepositoryFactory func() entity.RoomRepository
	InputPortFactory  func(repository entity.RoomRepository) port.RoomInputPort
}

func (c *RoomController) GetRoom(ctx *gin.Context) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	roomId := ctx.Param("roomId")
	i, err := strconv.Atoi(roomId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		return
	}
	res, err := inputport.GetRoomById(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(200, &res)
}

func (c *RoomController) GetRooms() (map[string]*port.RoomDto, error) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	rooms, err := inputport.GetRooms()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (c *RoomController) GetLobby(ctx *gin.Context) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	rooms, err := inputport.GetRooms()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		return
	}
	userName := ctx.GetString("userName")
	ctx.HTML(http.StatusOK, "lobby.tmpl", gin.H{
		"userName": userName,
		"rooms":    rooms,
	})
}

func (c *RoomController) InitRoom(ctx *gin.Context) {
	repository := c.RepositoryFactory()
	inputport := c.InputPortFactory(repository)
	roomId := ctx.Param("roomId")
	i, err := strconv.Atoi(roomId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		return
	}
	res, err := inputport.Init(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(200, &res)
}
