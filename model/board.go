package model

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"gopkg.in/olahol/melody.v1"
)

type User struct {
	Name string `json:"name"`
}

type Order struct {
	Type string `json:"type"`
}

type Player struct {
	User  `json:"user"`
	Order `json:"order"`
}

type Board struct {
	Id      int        `json:"id"`
	Pieces  [][]string `json:"pieces"`
	Turn    int        `json:"turn"`
	Status  string     `json:"status"`
	Players struct {
		First User `json:"first"`
		Draw  User `json:"draw"`
	} `json:"players"`
	Winner       string `json:"winner"`
	LastPutPoint string `json:"last_put_point"`
}

type Notify struct {
	Type  string
	Board Board
	Order string
}

func (b Board) GetBoardBinary() ([]byte, error) {
	url := "http://localhost:8000/api/v1/boards/" + strconv.Itoa(b.Id) + "/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byteArray, nil
}

func (b Board) GetBoard() (*Board, error) {
	bbyte, err := b.GetBoardBinary()
	if err != nil {
		return nil, err
	}
	var board Board
	if err := json.Unmarshal(bbyte, &board); err != nil {
		return nil, err
	}
	return &board, nil
}

func (b Board) PutPiece(point Point, s *melody.Session) ([]byte, error) {
	params := "?raw=" + strconv.Itoa(point.X) + "&column=" + strconv.Itoa(point.Y)
	url := "http://localhost:8000/api/v1/boards/" + strconv.Itoa(b.Id) + "/pieces/" + params

	player := new(Player)
	if userName, ok := s.Keys["userName"].(string); ok {
		user := User{Name: userName}
		player.User = user
	}
	if orderType, ok := s.Keys["order"].(string); ok {
		order := Order{Type: orderType}
		player.Order = order
	}
	playerJson, err := json.Marshal(player)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(playerJson))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byteArray, nil
}
