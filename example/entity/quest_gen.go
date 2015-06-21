package entity

import (
	"database/sql"

	"time"
)

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// Quest is generated quest table.
type Quest struct {
	ID       int        `db:"id"        goma:"size:11:pk"`
	Name     *string    `db:"name"      goma:""`
	Detail   *string    `db:"detail"    goma:""`
	CreateAt *time.Time `db:"create_at" goma:""`
}

// Scan Quest all scan
func (e *Quest) Scan(rows *sql.Rows) error {
	err := rows.Scan(&e.ID, &e.Name, &e.Detail, &e.CreateAt)
	return err
}
