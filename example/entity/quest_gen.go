package entity

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"time"
)

// QuestEntity is generated quest table.
type QuestEntity struct {
	ID       int
	Name     string
	Detail   string
	CreateAt time.Time
}
