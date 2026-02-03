package repository

import "devtrack/internal/model"


//define the interface with method for session repository ........
type SessionRepository interface {
	Create(session *model.Session) error
	GetActive() (*model.Session, error)

}
