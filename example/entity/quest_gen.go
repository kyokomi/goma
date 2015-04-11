package entity

import "database/sql"

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

// Scan QuestEntity all scan
func (e *QuestEntity) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.Name, &e.Detail, &e.CreateAt)
}
