package entity

import (
	"database/sql"
)

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaNumericTypes is generated goma_numeric_types table.
type GomaNumericTypes struct {
	ID              int64   `goma:":pk"`
	BoolColumns     bool    `goma:""`
	SmallintColumns int     `goma:""`
	IntColumns      int     `goma:""`
	IntegerColumns  int     `goma:""`
	SerialColumns   int     `goma:""`
	DecimalColumns  string  `goma:""`
	NumericColumns  string  `goma:""`
	FloatColumns    float64 `goma:""`
}

// Scan GomaNumericTypes all scan
func (e *GomaNumericTypes) Scan(rows *sql.Rows) error {
	err := rows.Scan(&e.ID, &e.BoolColumns, &e.SmallintColumns, &e.IntColumns, &e.IntegerColumns, &e.SerialColumns, &e.DecimalColumns, &e.NumericColumns, &e.FloatColumns)
	return err
}
