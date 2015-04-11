package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaNumericTypesEntity is generated goma_numeric_types table.
type GomaNumericTypesEntity struct {
	ID               int64   //`goma:"BIGINT(20):pk"`
	TinyintColumns   int     //`goma:"TINYINT(4)"`
	BoolColumns      int     //`goma:"TINYINT(1)"`
	SmallintColumns  int     //`goma:"SMALLINT(6)"`
	MediumintColumns int     //`goma:"MEDIUMINT(9)"`
	IntColumns       int     //`goma:"INT(11)"`
	IntegerColumns   int     //`goma:"INT(11)"`
	SerialColumns    int64   //`goma:"BIGINT(20)"`
	DecimalColumns   string  //`goma:"DECIMAL(10)"`
	NumericColumns   string  //`goma:"DECIMAL(10)"`
	FloatColumns     float32 //`goma:"FLOAT"`
	DoubleColumns    float64 //`goma:"DOUBLE"`
}

// Scan GomaNumericTypesEntity all scan
func (e *GomaNumericTypesEntity) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.TinyintColumns, &e.BoolColumns, &e.SmallintColumns, &e.MediumintColumns, &e.IntColumns, &e.IntegerColumns, &e.SerialColumns, &e.DecimalColumns, &e.NumericColumns, &e.FloatColumns, &e.DoubleColumns)
}
