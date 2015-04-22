package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"time"
)

// GomaDateTypes is generated goma_date_types table.
type GomaDateTypes struct {
	ID               int64     `migu:"size:20:pk"`
	DateColumns      time.Time `migu:""`
	DatetimeColumns  time.Time `migu:""`
	TimestampColumns time.Time `migu:""`
}

// Scan GomaDateTypes all scan
func (e *GomaDateTypes) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.DateColumns, &e.DatetimeColumns, &e.TimestampColumns)
}
