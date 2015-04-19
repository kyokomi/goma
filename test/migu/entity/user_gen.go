package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"time"
)

type User struct {
	ID       int64 `migu:"pk"`
	Name     string
	Email    *string
	Age      int
	CreateAt time.Time
	UpdateAt time.Time
}

// Scan User all scan
func (e *User) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.CreateAt, &e.UpdateAt)
}
