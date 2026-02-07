package service

import (
	"time"
	"errors"
"github.com/rahul-nakum14/devtrack/internal/model"
	"github.com/rahul-nakum14/devtrack/internal/repository"
)

type SessionService struct {
	repo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

// func (s *SessionService) StartSession(task, project string) (*model.Session, error) {
// 	session := &model.Session{
// 		Task:      task,
// 		Project:   project,
// 		StartTime: time.Now(),
// 	}

// 	if err := s.repo.Create(session); err != nil {
// 		return nil, err
// 	}

// 	return session, nil
// }

func (s *SessionService) StartSession(task, project string) (*model.Session, error) {
	active, err := s.GetActiveSession()
	if err != nil {
		return nil, err
	}

	if active != nil {
		return nil, errors.New(" session is  running")
	}

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

func (s *SessionService) GetActiveSession() (*model.Session,error){
	active,err := s.GetActiveSession()
	if err != nil {
		return nil, err
	}

	if active == nil {
		return nil, errors.New("no active session to stop")
	}

	return active,nil
}
func (s *SessionService) StopSession() (*model.Session, error) {
	active, err := s.GetActiveSession()
	if err != nil {
		return nil, err
	}

	if active == nil {
		return nil, errors.New("no active session to stop")
	}

	if err := s.repo.Stop(active); err != nil {
		return nil, err
	}

	return active, nil
}

func (s *SessionService) GetTodayStats() (map[string]time.Duration, time.Duration, error) {
	sessions, err := s.repo.GetTodaySessions()
	if err != nil {
		return nil, 0, err
	}

	perTask := make(map[string]time.Duration)
	var total time.Duration

	for _, session := range sessions {
		duration := session.EndTime.Sub(session.StartTime)
		perTask[session.Task] += duration
		total += duration
	}

	return perTask, total, nil
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
