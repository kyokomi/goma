package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"time"
)

type Quest struct {
	ID       int64  `migu:"pk"`
	Detail   string `migu:"size:512"`
	CreateAt time.Time
	UpdateAt time.Time
	Title    string
}

// Scan Quest all scan
func (e *Quest) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.Detail, &e.CreateAt, &e.UpdateAt, &e.Title)
}
