package routing

import (
	"encoding/json"
	"fmt"
	"go-gameroom/usecase/port"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type mInType struct {
	MsgType string `json:"type"`
}

type mJoinData struct {
	MsgData struct {
		Owner string `join:"owner"`
	} `json:"data"`
}

type mSelectData struct {
	MsgData struct {
		Order string `join:"order"`
	} `json:"data"`
}

type mPutPieceData struct {
	MsgData struct {
		PutPoint struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"putpoint"`
	} `json:"data"`
}

type mOutData struct {
	MsgType string `json:"Type"`
	MsgData struct {
		Pieces [10][10]string `json:"pieces"`
		Status string         `json:"status"`
		Turn   int            `json:"turn"`
	} `json:"data"`
}

type mOutBoardData struct {
	MsgType string `json:"Type"`
	MsgData struct {
		Pieces      [10][10]string `json:"pieces"`
		Owner       bool           `json:"owner"`
		Status      string         `json:"status"`
		FirstPlayer string         `json:"firstOrder"`
		DrawPlayer  string         `json:"drawOrder"`
		Order       string         `json:"order"`
		Turn        int            `json:"turn"`
	} `json:"data"`
}

func (r *GinRouter) UpgradeHandleFunc(c *gin.Context) {
	keys := map[string]interface{}{"userName": c.GetString("userName"), "roomId": c.Param("roomId")}
	r.melody.HandleRequestWithKeys(c.Writer, c.Request, keys)
}

func (r *GinRouter) setMelodyRouting() {
	r.melody.HandleMessage(func(s *melody.Session, msg []byte) {
		var mIntype mInType
		if err := json.Unmarshal(msg, &mIntype); err != nil {
			fmt.Println(err)
			return
		}
		switch mIntype.MsgType {
		case "join":
			roomId, _ := strconv.Atoi(s.MustGet("roomId").(string))
			userName := s.MustGet("userName").(string)
			var game *port.GetGameResponse
			game, _ = r.gameController.JoinGame(roomId, userName)

			var sendData mOutBoardData
			if userName == game.FirstPlayer {
				sendData.MsgData.Order = "first"
			} else if userName == game.DrawPlayer {
				sendData.MsgData.Order = "draw"
			}
			sendData.MsgType = "board-info"
			sendData.MsgData.Status = game.Status
			sendData.MsgData.FirstPlayer = game.FirstPlayer
			sendData.MsgData.DrawPlayer = game.DrawPlayer
			sendData.MsgData.Turn = game.Turn
			sendData.MsgData.Owner = false
			if game.Owner == userName {
				sendData.MsgData.Owner = true
			}
			bytes, _ := json.Marshal(sendData)
			s.Write(bytes)
		case "select-order":
			var receivedData mSelectData
			if err := json.Unmarshal(msg, &receivedData); err != nil {
				fmt.Println(err)
				return
			}
			roomId := s.MustGet("roomId").(string)
			roomIdInt, _ := strconv.Atoi(roomId)
			userName := s.MustGet("userName").(string)
			game, err := r.gameController.SelectGameOrder(roomIdInt, userName, receivedData.MsgData.Order)
			if err != nil {
				fmt.Println(err)
			}
			var sendData mOutBoardData
			sendData.MsgType = "board-info"
			sendData.MsgData.Status = game.Status
			sendData.MsgData.FirstPlayer = game.FirstPlayer
			sendData.MsgData.DrawPlayer = game.DrawPlayer
			sendData.MsgData.Turn = game.Turn
			if userName == game.FirstPlayer {
				sendData.MsgData.Order = "first"
			} else if userName == game.DrawPlayer {
				sendData.MsgData.Order = "draw"
			}
			bytes, err := json.Marshal(sendData)
			r.melody.Broadcast(bytes)

		case "put":
			var mPutPiecedata mPutPieceData
			if err := json.Unmarshal(msg, &mPutPiecedata); err != nil {
				fmt.Println(err)
				return
			}
			roomId := s.MustGet("roomId").(string)
			roomIdInt, _ := strconv.Atoi(roomId)
			userName := s.MustGet("userName").(string)
			game, err := r.gameController.PutPiece(mPutPiecedata.MsgData.PutPoint.X, mPutPiecedata.MsgData.PutPoint.Y, roomIdInt, userName)
			if err != nil {
				fmt.Println(err)
			}
			var sendData mOutData
			sendData.MsgType = "putresult"
			sendData.MsgData.Pieces = *game.Pieces
			sendData.MsgData.Status = game.Status
			sendData.MsgData.Turn = game.Turn
			bytes, err := json.Marshal(sendData)
			r.melody.Broadcast(bytes)
		}
	})
}
