package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaStringTypes is generated goma_string_types table.
type GomaStringTypes struct {
	ID                int64  `goma:"size:20:pk"`
	TextColumns       string `goma:""`
	TinytextColumns   string `goma:""`
	MediumtextColumns string `goma:""`
	LongtextColumns   string `goma:""`
	CharColumns       string `goma:"size:8"`
	VarcharColumns    string `goma:"size:255"`
}

// Scan GomaStringTypes all scan
func (e *GomaStringTypes) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.TextColumns, &e.TinytextColumns, &e.MediumtextColumns, &e.LongtextColumns, &e.CharColumns, &e.VarcharColumns)
}
