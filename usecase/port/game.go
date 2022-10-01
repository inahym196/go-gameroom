package port

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}

type GetGameResponse struct {
	Id     int
	Status string
}

type PutPieceResponse struct {
	Pieces *[10][10]string
}

type XOGameInputPort interface {
	GetGame(id int) (*GetGameResponse, error)
	GetGames() ([]*GetGameResponse, error)
	PutPiece(p *Point, id int) (*PutPieceResponse, error)
}
