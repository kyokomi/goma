package entity

import (
	"database/sql"

	"time"
)

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaDateTypes is generated goma_date_types table.
type GomaDateTypes struct {
	ID                  int64      `goma:":pk"`
	TimestampColumns    time.Time  `goma:""`
	NilTimestampColumns *time.Time `goma:""`
}

// Scan GomaDateTypes all scan
func (e *GomaDateTypes) Scan(rows *sql.Rows) error {
	err := rows.Scan(&e.ID, &e.TimestampColumns, &e.NilTimestampColumns)

	e.TimestampColumns = e.TimestampColumns.In(time.Local)

	if e.NilTimestampColumns != nil {
		_NilTimestampColumns := e.NilTimestampColumns.In(time.Local)
		e.NilTimestampColumns = &_NilTimestampColumns
	}

	return err
}
