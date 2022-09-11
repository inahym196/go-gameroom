package controller

import (
	"encoding/json"
	"errors"
	"go-gameroom/model"
	"log"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
	"gopkg.in/olahol/melody.v1"
)

func getNotify(b model.Board, s *melody.Session) ([]byte, error) {
	board, err := b.GetBoard()
	if err != nil {
		return nil, err
	}
	var notify model.Notify
	notify.Board = *board

	notify.Type = "notify"
	notifybyte, _ := json.Marshal(notify)
	return notifybyte, nil
}

func SetOrder(s *melody.Session, b model.Board) error {
	board, err := b.GetBoard()
	if err != nil {
		return err
	}
	var order string
	if s.Keys["userName"] == board.Players.First.Name {
		order = "first"
	} else if s.Keys["userName"] == board.Players.Draw.Name {
		order = "draw"
	} else {
		order = "audience"
	}
	s.Set("order", order)
	return nil
}

type ClientMsg []byte

func (msg ClientMsg) parseType() (string, error) {
	value := gjson.Get(string(msg), "type")
	if !value.Exists() {
		return "", errors.New("not exists")
	}
	return value.String(), nil
}

func (msg ClientMsg) parsePoint() (*model.Point, error) {
	pieceMap := gjson.Get(string(msg), "piece")
	if !pieceMap.Exists() {
		return nil, errors.New("not exists")
	}
	pieceByte, err := json.Marshal(pieceMap.Value())
	if err != nil {
		return nil, err
	}
	var point model.Point
	if err := json.Unmarshal(pieceByte, &point); err != nil {
		return nil, err
	}
	return &point, nil
}

func GetMRouter() *melody.Melody {
	m := melody.New()
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		var b model.Board
		boardId, err := strconv.Atoi(strings.Split(s.Request.RequestURI, "/")[2])
		if err != nil {
			log.Fatal(err)
		}
		b.Id = boardId
		clientMsg := ClientMsg(msg)
		msgType, err := clientMsg.parseType()
		if err != nil {
			log.Fatal(err)
		}
		switch msgType {
		case "join":
			SetOrder(s, b)
			board, err := b.GetBoard()
			if err != nil {
				log.Fatal(err)
			}
			var notify model.Notify
			notify.Board = *board
			if order, ok := s.Keys["order"].(string); ok {
				notify.Order = order
			}
			notify.Type = "join"
			notifyByte, err := json.Marshal(notify)
			if err != nil {
				log.Fatal(err)
			}
			s.Write(notifyByte)
		case "put":
			point, err := clientMsg.parsePoint()
			if err != nil {
				log.Fatal(err)
			}
			byteArray, err := b.PutPiece(*point, s)
			if err != nil {
				log.Fatal(err)
			}
			m.Broadcast(byteArray)
		}
	})
	return m
}
