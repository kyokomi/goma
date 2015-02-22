package entity

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaBinaryTypesEntity is generated goma_binary_types table.
type GomaBinaryTypesEntity struct {
	ID                int64   //`goma:"BIGINT(20):pk"`
	BinaryColumns     []uint8 //`goma:"BINARY(3)"`
	TinyblobColumns   []uint8 //`goma:"TINYBLOB"`
	BlobColumns       []uint8 //`goma:"BLOB"`
	MediumblobColumns []uint8 //`goma:"MEDIUMBLOB"`
	LongblobColumns   []uint8 //`goma:"LONGBLOB"`
	VarbinaryColumns  []uint8 //`goma:"VARBINARY(10)"`
}
