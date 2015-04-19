package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaStringTypesEntity is generated goma_string_types table.
type GomaStringTypesEntity struct {
	ID             int64  `goma:"BIGINT:pk"`
	TextColumns    string `goma:"TEXT"`
	CharColumns    string `goma:"VARCHAR"`
	VarcharColumns string `goma:"VARCHAR"`
}

// Scan GomaStringTypesEntity all scan
func (e *GomaStringTypesEntity) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.TextColumns, &e.CharColumns, &e.VarcharColumns)
}
