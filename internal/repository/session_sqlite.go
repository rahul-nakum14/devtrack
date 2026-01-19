package repository

import (
	"database/sql"
	"time"

	"devtrack/internal/model"
)

type SessionSQLiteRepository struct {
	db *sql.DB
}

func NewSessionSQLiteRepository(db *sql.DB) *SessionSQLiteRepository {
	return &SessionSQLiteRepository{db: db}
}

func (r *SessionSQLiteRepository) Create(session *model.Session) error {
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

	_, err := r.db.Exec(
		`INSERT INTO sessions (task, project, start_time) VALUES (?, ?, ?)`,
		session.Task,
		session.Project,
		session.StartTime,
	)
	return err
}

// yet to implement active and stop methsds