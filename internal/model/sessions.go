package model

import "time"

type Session struct {
	ID        int
	Task      string
	Project   string
	StartTime time.Time
	EndTime   *time.Time
}
