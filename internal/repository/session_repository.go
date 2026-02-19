package repository

import "github.com/rahul-nakum14/devtrack/internal/model"


//define the interface with method for session repository ........
type SessionRepository interface {
	Create(session *model.Session) error
	GetActiveSession() (*model.Session, error)
	Stop(session *model.Session) error
	GetTodaySessions() ([]*model.Session, error)
	GetWeekSessions() ([]*model.Session, error)
	Migrate()(error)
}
