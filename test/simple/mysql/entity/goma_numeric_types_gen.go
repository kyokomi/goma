package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaNumericTypes is generated goma_numeric_types table.
type GomaNumericTypes struct {
	ID               int64   `migu:"size:20:pk"`
	TinyintColumns   int     `migu:"size:4"`
	BoolColumns      int     `migu:"size:1"`
	SmallintColumns  int     `migu:"size:6"`
	MediumintColumns int     `migu:"size:9"`
	IntColumns       int     `migu:"size:11"`
	IntegerColumns   int     `migu:"size:11"`
	SerialColumns    int64   `migu:"size:20"`
	DecimalColumns   string  `migu:"size:10"`
	NumericColumns   string  `migu:"size:10"`
	FloatColumns     float32 `migu:""`
	DoubleColumns    float64 `migu:""`
}

// Scan GomaNumericTypes all scan
func (e *GomaNumericTypes) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.TinyintColumns, &e.BoolColumns, &e.SmallintColumns, &e.MediumintColumns, &e.IntColumns, &e.IntegerColumns, &e.SerialColumns, &e.DecimalColumns, &e.NumericColumns, &e.FloatColumns, &e.DoubleColumns)
}
