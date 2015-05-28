package entity

import (
	"database/sql"
)

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaBinaryTypes is generated goma_binary_types table.
type GomaBinaryTypes struct {
	BinaryID          int64   `goma:"size:20:pk"`
	BinaryColumns     []uint8 `goma:"size:3"`
	TinyblobColumns   []uint8 `goma:""`
	BlobColumns       []uint8 `goma:""`
	MediumblobColumns []uint8 `goma:""`
	LongblobColumns   []uint8 `goma:""`
	VarbinaryColumns  []uint8 `goma:"size:10"`
}

// Scan GomaBinaryTypes all scan
func (e *GomaBinaryTypes) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.BinaryID, &e.BinaryColumns, &e.TinyblobColumns, &e.BlobColumns, &e.MediumblobColumns, &e.LongblobColumns, &e.VarbinaryColumns)
}
