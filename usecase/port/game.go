package port

import (
	"go-gameroom/domain/entity"
)

/*
 * Input Port
 *  └─ Interactor で実装、Controller で使用される
 */
type GameInputPort interface {
	GetGames()
	GetGame(*GetGameRequestParams)
}

type GetGameRequestParams struct {
	GameId string
}

/*
 * Output Port
 *  └─ Presenter で実装、Interactor で使用される
 */
type GameOutputPort interface {
	GetGames([]entity.Game)
	GetGame(*entity.Game)
}

type GetGameResponse struct {
	Game *entity.Game
}

type GetGamesResponse struct {
	Games []entity.Game
}
