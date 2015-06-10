package entity

import (
	"database/sql"
	"database/sql/driver"
)

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaStringTypes is generated goma_string_types table.
type GomaStringTypes struct {
	ID                int64       `goma:"size:20:pk"`
	TextColumns       string      `goma:""`
	TinytextColumns   string      `goma:""`
	MediumtextColumns string      `goma:""`
	LongtextColumns   string      `goma:""`
	CharColumns       string      `goma:"size:8"`
	VarcharColumns    string      `goma:"size:255"`
	EnumColumns       EnumColumns `goma:""`
}

// Scan GomaStringTypes all scan
func (e *GomaStringTypes) Scan(rows *sql.Rows) error {
	err := rows.Scan(&e.ID, &e.TextColumns, &e.TinytextColumns, &e.MediumtextColumns, &e.LongtextColumns, &e.CharColumns, &e.VarcharColumns, &e.EnumColumns)
	return err
}

// EnumColumns EnumColumns column type
type EnumColumns string

// EnumColumns
const (
	EnumColumnsOpen    EnumColumns = "OPEN"
	EnumColumnsClose   EnumColumns = "CLOSE"
	EnumColumnsForever EnumColumns = "FOREVER"
)

// Scan database/sql Scanner
func (e *EnumColumns) Scan(v interface{}) error {
	*e = EnumColumns(v.([]byte))
	return nil
}

var _ sql.Scanner = (*EnumColumns)(nil)

// Value database/sql/driver Valuer
func (e EnumColumns) Value() (driver.Value, error) {
	return string(e), nil
}

var _ driver.Valuer = (*EnumColumns)(nil)
