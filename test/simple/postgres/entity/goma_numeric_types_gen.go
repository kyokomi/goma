package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaNumericTypes is generated goma_numeric_types table.
type GomaNumericTypes struct {
	ID              int64   `migu:":pk"`
	BoolColumns     bool    `migu:""`
	SmallintColumns int     `migu:""`
	IntColumns      int     `migu:""`
	IntegerColumns  int     `migu:""`
	SerialColumns   int     `migu:""`
	DecimalColumns  string  `migu:""`
	NumericColumns  string  `migu:""`
	FloatColumns    float64 `migu:""`
}

// Scan GomaNumericTypes all scan
func (e *GomaNumericTypes) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.BoolColumns, &e.SmallintColumns, &e.IntColumns, &e.IntegerColumns, &e.SerialColumns, &e.DecimalColumns, &e.NumericColumns, &e.FloatColumns)
}
