package gateway

import (
	"go-gameroom/domain/entity"
)

type GameRedisRepository struct {
}

func NewGameRedisRepository() entity.GameRepository {
	return &GameRedisRepository{}
}

func (repo *GameRedisRepository) Find(id entity.GameId) (*entity.Game, error) {
	/* GameIdからGameを取得する具体的な実装 */
	game := entity.Game{Id: id}
	return &game, nil
}

func (repo *GameRedisRepository) FindAll() ([]entity.Game, error) {
	/* Gamesを取得する具体的な実装 */
	var Games []entity.Game
	return Games, nil
}

func (repo *GameRedisRepository) Regist(*entity.Game) error {
	/* Gameを登録する具体的な実装 */
	return nil
}
