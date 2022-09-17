package entity

type GameStatus int

const (
	Init GameStatus = iota
	Waiting
	Starting
	End
)

type Game struct {
	Status  GameStatus
	Players []User
	Winner  *User
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
