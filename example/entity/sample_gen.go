package entity

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"time"
)

// SampleEntity is generated sample table.
type SampleEntity struct {
	ID       int       //`goma:"INT(11):pk"`
	Name     string    //`goma:"TEXT"`
	CreateAt time.Time //`goma:"TIMESTAMP"`
}
