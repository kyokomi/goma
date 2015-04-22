package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"time"
)

// Quest is generated quest table.
type Quest struct {
	ID       int64     `migu:"size:20:pk"`
	Title    string    `migu:"size:255"`
	Detail   string    `migu:"size:512"`
	CreateAt time.Time `migu:""`
	UpdateAt time.Time `migu:""`
}

// Scan Quest all scan
func (e *Quest) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.Title, &e.Detail, &e.CreateAt, &e.UpdateAt)
}
