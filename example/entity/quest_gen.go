package entity

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"time"
)

// QuestEntity is generated quest table.
type QuestEntity struct {
	ID       int       //`goma:"INT(11):pk"`
	Name     string    //`goma:"TEXT"`
	Detail   string    //`goma:"TEXT"`
	CreateAt time.Time //`goma:"DATETIME"`
}
