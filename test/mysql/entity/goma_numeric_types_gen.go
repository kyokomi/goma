package entity

import (
	"database/sql"
)

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaNumericTypes is generated goma_numeric_types table.
type GomaNumericTypes struct {
	ID               int64   `goma:"size:20:pk"`
	TinyintColumns   int     `goma:"size:4"`
	BoolColumns      bool    `goma:"size:1"`
	SmallintColumns  int     `goma:"size:6"`
	MediumintColumns int     `goma:"size:9"`
	IntColumns       int     `goma:"size:11"`
	IntegerColumns   int     `goma:"size:11"`
	SerialColumns    int64   `goma:"size:20"`
	DecimalColumns   string  `goma:"size:10"`
	NumericColumns   string  `goma:"size:10"`
	FloatColumns     float32 `goma:""`
	DoubleColumns    float64 `goma:""`
}

// Scan GomaNumericTypes all scan
func (e *GomaNumericTypes) Scan(rows *sql.Rows) error {
	err := rows.Scan(&e.ID, &e.TinyintColumns, &e.BoolColumns, &e.SmallintColumns, &e.MediumintColumns, &e.IntColumns, &e.IntegerColumns, &e.SerialColumns, &e.DecimalColumns, &e.NumericColumns, &e.FloatColumns, &e.DoubleColumns)

	return err
}
