package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaBinaryTypes is generated goma_binary_types table.
type GomaBinaryTypes struct {
	BinaryID          int64   `migu:"size:20:pk"`
	BinaryColumns     []uint8 `migu:"size:3"`
	TinyblobColumns   []uint8 `migu:""`
	BlobColumns       []uint8 `migu:""`
	MediumblobColumns []uint8 `migu:""`
	LongblobColumns   []uint8 `migu:""`
	VarbinaryColumns  []uint8 `migu:"size:10"`
}

// Scan GomaBinaryTypes all scan
func (e *GomaBinaryTypes) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.BinaryID, &e.BinaryColumns, &e.TinyblobColumns, &e.BlobColumns, &e.MediumblobColumns, &e.LongblobColumns, &e.VarbinaryColumns)
}
