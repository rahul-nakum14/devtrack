package repository

import (
	"database/sql"
	"time"

	"github.com/rahul-nakum14/devtrack/internal/model"
)

type SessionSQLiteRepository struct {
	db *sql.DB
}

func NewSessionSQLiteRepository(db *sql.DB) *SessionSQLiteRepository {
	return &SessionSQLiteRepository{db: db}
}


func (r * SessionSQLiteRepository) Migrate() error{
	query := `
	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL,
		project TEXT,
		start_time DATETIME NOT NULL,
		end_time DATETIME
	);
	`
	if _, err := r.db.Exec(query); err != nil {
	
		return err
	}

	return nil
}

func (r *SessionSQLiteRepository) Create(session *model.Session) error {
	_, err := r.db.Exec(
		`INSERT INTO sessions (task, project, start_time) VALUES (?, ?, ?)`,
		session.Task,
		session.Project,
		session.StartTime,
	)
	
	return err
}

func (r *SessionSQLiteRepository) GetActiveSession() (*model.Session, error){	
	row := r.db.QueryRow(`
		SELECT id, task, project, start_time
		FROM sessions
		WHERE end_time IS NULL
		ORDER BY start_time DESC
		LIMIT 1
	`)
	var session model.Session

	err:= row.Scan(
		&session.ID,
		&session.Task,
		&session.Project,
		&session.StartTime,
	)
	if err == sql.ErrNoRows {
	
		return nil, nil
	}
	if err != nil{
		return nil, err
	}

	return &session, nil;
}

func (r *SessionSQLiteRepository) Stop(session *model.Session) error {
	now := time.Now()
	session.EndTime = &now

	_, err := r.db.Exec(
		`UPDATE sessions SET end_time = ? WHERE id = ?`,
		now,
		session.ID,
	)

	return err
}

func (r *SessionSQLiteRepository) GetTodaySessions() ([]*model.Session, error) {
	rows, err := r.db.Query(`
		SELECT id, task, project, start_time, end_time
		FROM sessions
		WHERE date(start_time) = date('now', 'localtime')
		  AND end_time IS NOT NULL
	`)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	var sessions []*model.Session

	for rows.Next() {
		var s model.Session
		var endTime time.Time

		err := rows.Scan(
			&s.ID,
			&s.Task,
			&s.Project,
			&s.StartTime,
			&endTime,
		)
		if err != nil {
			return nil, err
		}

		s.EndTime = &endTime
		sessions = append(sessions, &s)
	}

	return sessions, nil
}

func (r *SessionSQLiteRepository) GetWeekSessions() ([]*model.Session, error) {

	rows, err := r.db.Query(`
		SELECT id, task, project, start_time, end_time
		FROM sessions
		WHERE start_time >= datetime('now', '-7 days')
		AND end_time IS NOT NULL
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []*model.Session

	for rows.Next() {
		var session model.Session

		err := rows.Scan(
			&session.ID,
			&session.Task,
			&session.Project,
			&session.StartTime,
			&session.EndTime,
		)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, &session)
	}

	return sessions, nil
}

func (r *SessionSQLiteRepository) GetAllSessions() ([]*model.Session, error) {

	rows, err := r.db.Query(`
		SELECT id, task, project, start_time, end_time
	     FROM sessions
		ORDER BY start_time DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []*model.Session

	for rows.Next() {
		var session model.Session

		err := rows.Scan(
			&session.ID,
			&session.Task,
			&session.Project,
			&session.StartTime,
			&session.EndTime,
		)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, &session)
	}

	return sessions, nil
}


// yet to implement active and stop methsds