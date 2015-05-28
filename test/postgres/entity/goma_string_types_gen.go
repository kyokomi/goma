package entity

import (
	"database/sql"
)

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaStringTypes is generated goma_string_types table.
type GomaStringTypes struct {
	ID             int64  `goma:":pk"`
	TextColumns    string `goma:""`
	CharColumns    string `goma:""`
	VarcharColumns string `goma:""`
}

// Scan GomaStringTypes all scan
func (e *GomaStringTypes) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.TextColumns, &e.CharColumns, &e.VarcharColumns)
}
