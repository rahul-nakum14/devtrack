package service

import (
	"fmt"
	"time"

	"devtrack/internal/model"
)

type SessionService struct{}

func NewSessionService() *SessionService {
	return &SessionService{}
}

func (s *SessionService) StartSession(task, project string) *model.Session {
	session := &model.Session{
		Task:      task,
		Project:   project,
		StartTime: time.Now(),
	}

	fmt.Println("Session started:")
	fmt.Println(" Task   :", session.Task)

	if session.Project != "" {
		fmt.Println(" Project:", session.Project)
	}

	fmt.Println(" Started:", session.StartTime.Format(time.RFC3339))

	return session
}
