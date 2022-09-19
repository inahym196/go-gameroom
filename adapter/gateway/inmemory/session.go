package gateway

import (
	"fmt"
	"go-gameroom/domain/entity"
	"go-gameroom/usecase/repository"
	"time"

	"github.com/google/uuid"
)

type SessionRepository struct{}

var SessionDataBase map[string]*entity.Session = (make(map[string]*entity.Session))

func NewSessionRepository() repository.SessionRepository {
	return &SessionRepository{}
}

func (s *SessionRepository) Get(id string) (*entity.Session, error) {
	if session, ok := SessionDataBase[id]; ok {
		return session, nil
	}
	return nil, fmt.Errorf("Session Not Found")
}

func (s *SessionRepository) Create(data string) (*entity.Session, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	session := &entity.Session{Id: uuid.String(), Data: data, UpdatedAt: time.Now().String()}
	SessionDataBase[session.Id] = session
	return session, nil
}
