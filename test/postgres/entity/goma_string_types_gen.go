package entity

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

// GomaStringTypesEntity is generated goma_string_types table.
type GomaStringTypesEntity struct {
	ID             int64  //`goma:"BIGINT:pk"`
	TextColumns    string //`goma:"TEXT"`
	CharColumns    string //`goma:"VARCHAR"`
	VarcharColumns string //`goma:"VARCHAR"`
}
