package model

import (
	"encoding/json"
	"strconv"
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

func GetBoard(boardId int) (*Board, error) {
	bbyte, err := GetBoardBinary(boardId)
	if err != nil {
		return nil, err
	}
	var board Board
	if err := json.Unmarshal(bbyte, &board); err != nil {
		return nil, err
	}
	return &board, nil
}

func GetBoardBinary(boardId int) ([]byte, error) {
	var req Request
	req.url = "http://localhost:8000/api/v1/boards/" + strconv.Itoa(boardId) + "/"
	return req.Get()
}
