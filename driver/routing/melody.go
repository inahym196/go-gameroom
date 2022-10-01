package routing

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type mInData struct {
	MsgType string `json:"type"`
	MsgData struct {
		PutPoint struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"putpoint"`
		GameType string `json:"gameType"`
	} `json:"data"`
}

type mOutData struct {
	MsgType string `json:"type"`
	MsgData struct {
		Pieces [10][10]string `json:"pieces"`
		Status string         `json:"status"`
	} `json:"data"`
}

func (r *GinRouter) UpgradeHandleFunc(c *gin.Context) {
	keys := map[string]interface{}{"userName": c.GetString("userName"), "roomId": c.Param("roomId")}
	r.melody.HandleRequestWithKeys(c.Writer, c.Request, keys)
}

func (r *GinRouter) setMelodyRouting() {
	r.melody.HandleMessage(func(s *melody.Session, msg []byte) {
		var mIndata mInData
		if err := json.Unmarshal(msg, &mIndata); err != nil {
			fmt.Println(err)
			return
		}
		switch mIndata.MsgType {
		case "join":
			gameType := mIndata.MsgData.GameType
			if gameType != "XOGame" {
				return
			}
			roomId := s.MustGet("roomId").(string)
			roomIdInt, _ := strconv.Atoi(roomId)
			game, err := r.gameController.GetGame(roomIdInt)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("game: %#v\n", game)
			var mOutdata mOutData
			mOutdata.MsgType = "board-info"
			mOutdata.MsgData.Status = game.Status
			bytes, err := json.Marshal(mOutdata)
			s.Write(bytes)
			// fmt.Printf("join: %#v\n", mdata)
		case "put":
			roomId := s.MustGet("roomId").(string)
			roomIdInt, _ := strconv.Atoi(roomId)
			res, err := r.gameController.PutPiece(mIndata.MsgData.PutPoint.X, mIndata.MsgData.PutPoint.Y, "X", roomIdInt)
			if err != nil {
				fmt.Println(err)
			}
			for _, col := range res.Pieces {
				fmt.Printf("  %v\n", col)
			}
			var mOutdata mOutData
			mOutdata.MsgType = "putresult"
			mOutdata.MsgData.Pieces = *res.Pieces
			bytes, err := json.Marshal(mOutdata)
			s.Write(bytes)
		}
	})
}
