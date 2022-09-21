package entity

type GameStatus string

const (
	Init     = GameStatus("Init")
	Waiting  = GameStatus("Waiting")
	Starting = GameStatus("Starting")
	End      = GameStatus("End")
)

type Game struct {
	status  GameStatus
	players []User
	winner  *User
}

func (g *Game) Status() string {
	return string(Init)
}

type PutPoint struct {
	x int
	y int
}

type SquareBoardGame struct {
	Game
	Turn         int
	Pieces       [][]string
	LastPutPoint PutPoint
}

type XOGame struct {
	SquareBoardGame
}

type GameRepository interface {
	Init() error
	Join(*User) bool
}

type XOGameRepository interface {
	GameRepository
	Get() (*XOGame, error)
	PutPiece(*PutPoint) (*XOGame, error)
}

func NewXOGame() *XOGame {
	return &XOGame{}
}
