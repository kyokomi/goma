package entity

import (
	"database/sql"

	"time"
)

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// Sample is generated sample table.
type Sample struct {
	ID       int       `goma:"size:11:pk"`
	Name     *string   `goma:""`
	CreateAt time.Time `goma:""`
}

// Scan Sample all scan
func (e *Sample) Scan(rows *sql.Rows) error {
	err := rows.Scan(&e.ID, &e.Name, &e.CreateAt)
	e.CreateAt = e.CreateAt.In(time.Local)
	return err
}
