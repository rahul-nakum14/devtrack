package service

import (
	"time"

	"devtrack/internal/model"
	"devtrack/internal/repository"
)

type SessionService struct {
	repo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

func (s *SessionService) StartSession(task, project string) (*model.Session, error) {
	session := &model.Session{
		Task:      task,
		Project:   project,
		StartTime: time.Now(),
	}

	if err := s.repo.Create(session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *SessionService) GetActiveSession() (*model.Session, error) {
	session, err := s.repo.GetActive()
	if err != nil {
		return nil, err
	}	

	return session, nil
}


// package service

// import (
// 	"fmt"
// 	"time"

// 	"devtrack/internal/model"
// )

// type SessionService struct{}

// func NewSessionService() *SessionService {
// 	return &SessionService{}
// }

// func (s *SessionService) StartSession(task, project string) *model.Session {
// 	session := &model.Session{
// 		Task:      task,
// 		Project:   project,
// 		StartTime: time.Now(),
// 	}

// 	fmt.Println("Session started:")
// 	fmt.Println(" Task   :", session.Task)

// 	if session.Project != "" {
// 		fmt.Println(" Project:", session.Project)
// 	}

// 	fmt.Println(" Started:", session.StartTime.Format(time.RFC3339))

// 	return session
// }
