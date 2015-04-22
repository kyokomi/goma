package models

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// Scan User all scan
func (e *User) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.CreateAt, &e.UpdateAt)
}
