package contracts

import "github.com/ladmakhi81/gobanks/entities"

type SessionRepository interface {
	CreateSession(session *entities.Session) error
	DeleteSessionByUserId(userID int) error
	GetSessionByToken(accessToken string) (*entities.Session, error)
}
