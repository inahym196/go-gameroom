package port

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}

type GetGameResponse struct {
	Id          int
	Owner       string
	Status      string
	FirstPlayer string
	DrawPlayer  string
	Turn        int
}

type PutPieceResponse struct {
	Pieces *[10][10]string
	Status string
	Turn   int
}

type XOGameInputPort interface {
	GetGame(id int) (*GetGameResponse, error)
	GetGames() ([]*GetGameResponse, error)
	SelectGameOrder(id int, user string, order string) (*GetGameResponse, error)
	InitGame(id int, owner string) (*GetGameResponse, error)
	JoinGame(id int, user string) (*GetGameResponse, error)
	PutPiece(id int, p *Point, user string) (*PutPieceResponse, error)
}
