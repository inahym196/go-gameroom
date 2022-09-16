package entity

type GameId string

type Game struct {
	Id GameId
}

type GameRepository interface {
	Find(GameId) (*Game, error)
	FindAll() ([]Game, error)
	Regist(*Game) error
}
