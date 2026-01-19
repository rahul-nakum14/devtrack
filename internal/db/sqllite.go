package db

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() (*sql.DB, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dbDir := filepath.Join(home, ".devtrack")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, err
	}

	dbPath := filepath.Join(dbDir, "devtrack.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}
