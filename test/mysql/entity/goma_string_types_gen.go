package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaStringTypesEntity is generated goma_string_types table.
type GomaStringTypesEntity struct {
	ID                int64  `goma:"BIGINT(20):pk"`
	TextColumns       string `goma:"TEXT"`
	TinytextColumns   string `goma:"TINYTEXT"`
	MediumtextColumns string `goma:"MEDIUMTEXT"`
	LongtextColumns   string `goma:"LONGTEXT"`
	CharColumns       string `goma:"CHAR(8)"`
	VarcharColumns    string `goma:"VARCHAR(255)"`
}

// Scan GomaStringTypesEntity all scan
func (e *GomaStringTypesEntity) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.TextColumns, &e.TinytextColumns, &e.MediumtextColumns, &e.LongtextColumns, &e.CharColumns, &e.VarcharColumns)
}
