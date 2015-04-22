package models

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// Scan Quest all scan
func (e *Quest) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.Title, &e.Detail, &e.CreateAt, &e.UpdateAt)
}
