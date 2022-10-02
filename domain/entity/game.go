package entity

import (
	"fmt"
)

type gameStatus string

const (
	gameInit     = gameStatus("Init")
	gameSetting  = gameStatus("Setting")
	gameWaiting  = gameStatus("Waiting")
	gameStarting = gameStatus("Starting")
	gameEnd      = gameStatus("End")
)

type Game struct {
	id     int
	status gameStatus
	owner  string
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

func (g *Game) GetOwner() string {
	return g.owner
}

type Point struct {
	X, Y int
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
	players      XOPlayers
	winner       string
	turn         int
	lastPutPoint Point
	pieces       *[10][10]XOPiece
}

type XOPlayers struct {
	first string
	draw  string
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
	xoGame.status = gameInit
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
		xg.SetNextTurn()
	default:
		return nil, fmt.Errorf("Invalid piece type")
	}
	return xg.pieces, nil
}

func (xg *XOGame) GetPlayers() (string, string) {
	return string(xg.players.first), string(xg.players.draw)
}

func (xg *XOGame) SetPlayer(user string, order string) error {
	if xg.owner == "" && xg.owner != user {
		return fmt.Errorf("owner not exists")
	}
	if order == "first" {
		if xg.players.first != "" {
			return fmt.Errorf("already exists first")
		}
		xg.players.first = user
		xg.updateStatus()
	} else if order == "draw" {
		if xg.players.draw != "" {
			return fmt.Errorf("already esists draw")
		}
		xg.players.draw = user
		xg.updateStatus()
	} else {
		return fmt.Errorf("order syntax errror")
	}
	return nil
}

func (xg *XOGame) updateStatus() {
	players := xg.players
	if xg.winner != "" {
		fmt.Println("status: gameEnd")
		xg.status = gameEnd
	} else if xg.owner == "" {
		xg.status = gameInit
		fmt.Println("status: gameInit")
	} else if players.first == "" && players.draw == "" {
		xg.status = gameSetting
		fmt.Println("status: gameSetting")
	} else if players.first == "" || players.draw == "" {
		xg.status = gameWaiting
		fmt.Println("status: gameWaiting")
	} else if players.first != "" && players.draw != "" {
		xg.status = gameStarting
		fmt.Println("status: gameStarting")
	}
}

func (xg *XOGame) BecomeOwner(user string) error {
	if xg.owner != "" {
		return fmt.Errorf("owner already exists")
	}
	xg.owner = user
	xg.updateStatus()
	return nil
}

func (xg *XOGame) GetOrder(user string) string {
	players := xg.players
	if players.first == user {
		return "first"
	} else if players.draw == user {
		return "draw"
	}
	return ""
}

func (xg *XOGame) GetPlayerByOrder(order string) string {
	if order == "first" {
		return xg.players.first
	} else if order == "draw" {
		return xg.players.draw
	}
	return ""
}

func (xg *XOGame) GetTurn() int {
	return xg.turn
}

func (xg *XOGame) SetNextTurn() {
	xg.turn = xg.turn + 1
}

func (xg *XOGame) GetNextPiece() string {
	turn := xg.GetTurn()
	var piece XOPiece
	if turn%2 == 0 {
		if turn%8 == 6 {
			piece = PieceOG
		} else {
			piece = PieceOP
		}
	} else {
		if turn%8 == 7 {
			piece = PieceXP
		} else {
			piece = PieceXG
		}
	}
	return string(piece)
}
