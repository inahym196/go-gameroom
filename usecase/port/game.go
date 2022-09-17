package port

type GameId = string
type GameStatus = int

const (
	Init GameStatus = iota
	Waiting
	Starting
	End
)

type Game struct {
	Status GameStatus
}

type XOGameInputPort interface {
	GetGame()
}

type XOGame struct {
	Game
}

type XOGameOutputPort interface {
	GetGame(*XOGame)
}

type GetXOGameResponse struct {
	Game *XOGame
}
