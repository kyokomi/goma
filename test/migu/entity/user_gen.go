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
	CreateAt time.Time
	UpdateAt time.Time
	Age      int
}

// Scan User all scan
func (e *User) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.Name, &e.Email, &e.CreateAt, &e.UpdateAt, &e.Age)
}
