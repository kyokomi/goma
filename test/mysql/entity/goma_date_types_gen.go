package entity

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"time"
)

// GomaDateTypesEntity is generated goma_date_types table.
type GomaDateTypesEntity struct {
	ID               int64
	DateColumns      time.Time
	DatetimeColumns  time.Time
	TimestampColumns time.Time
}
