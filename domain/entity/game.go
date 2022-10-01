package entity

import (
	"fmt"
)

type gameStatus string

const (
	gameInit     = gameStatus("Init")
	gameWaiting  = gameStatus("Waiting")
	gameStarting = gameStatus("Starting")
	gameEnd      = gameStatus("End")
)

type Game struct {
	id     int
	status gameStatus
}

func (g *Game) GetId() int {
	return g.id
}

func (g *Game) GetStatus() string {
	return string(g.status)
}

func (g *Game) SetStatus(status string) error {
	setStatus := gameStatus(status)
	switch setStatus {
	case gameInit, gameWaiting, gameStarting, gameEnd:
		g.status = setStatus
		return nil
	}
	return fmt.Errorf("Status is invalid")
}

type Point struct {
	X, Y int
}

type XOPlayer struct {
	order string
	user  *User
}

type XOPiece string

const (
	PieceNone = XOPiece("None")
	PieceXP   = XOPiece("XP")
	PieceXG   = XOPiece("XG")
	PieceOP   = XOPiece("OP")
	PieceOG   = XOPiece("OG")
)

func (xp *XOPiece) String() string {
	return string(*xp)
}

type XOGame struct {
	Game
	players      [2]*XOPlayer
	winner       *XOPlayer
	turn         int
	lastPutPoint Point
	pieces       *[10][10]XOPiece
}

func NewXOGame(id int) *XOGame {
	var pieces [10][10]XOPiece
	for i, col := range pieces {
		for j := range col {
			pieces[i][j] = PieceNone
		}
	}
	xoGame := &XOGame{pieces: &pieces}
	xoGame.id = id
	xoGame.status = "Waiting"
	return xoGame
}

func (xg *XOGame) GetPieces() *[10][10]XOPiece {
	return xg.pieces
}

func (xg *XOGame) SetPiece(x, y int, piece string) (*[10][10]XOPiece, error) {
	if xg.pieces[y][x] != PieceNone {
		return nil, fmt.Errorf("There is already exists in this location")
	}
	setPiece := XOPiece(piece)
	switch setPiece {
	case PieceOG, PieceOP, PieceXG, PieceXP:
		xg.pieces[y][x] = setPiece
	default:
		return nil, fmt.Errorf("Invalid piece type")
	}
	return xg.pieces, nil
}

type Player struct {
	user  *User
	order string
}

func NewPlayer(userName string) *Player {
	return &Player{
		user: &User{},
	}
}
