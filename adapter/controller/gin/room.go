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
	PortFactory       func(repository entity.RoomRepository) port.RoomLoosePort
}

func (c *RoomController) GetRoom(ctx *gin.Context) {
	repository := c.RepositoryFactory()
	port := c.PortFactory(repository)
	roomId := ctx.Param("roomId")
	i, err := strconv.Atoi(roomId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		return
	}
	res, err := port.GetRoomById(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(200, &res)
}

func (c *RoomController) GetRooms(ctx *gin.Context) {
	repository := c.RepositoryFactory()
	port := c.PortFactory(repository)
	res, err := port.GetRooms()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(200, &res)
}

func (c *RoomController) InitRoom(ctx *gin.Context) {
	repository := c.RepositoryFactory()
	port := c.PortFactory(repository)
	roomId := ctx.Param("roomId")
	i, err := strconv.Atoi(roomId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		return
	}
	res, err := port.Init(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(200, &res)
}
