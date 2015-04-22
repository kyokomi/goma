package entity

import "database/sql"

// NOTE: THIS FILE WAS PRODUCED BY THE
// GOMA CODE GENERATION TOOL (github.com/kyokomi/goma)
// DO NOT EDIT

import (
	"time"
)

// User is generated user table.
type User struct {
	ID       int64     `migu:"size:20:pk"`
	Name     string    `migu:"size:255"`
	Email    string    `migu:"size:255"`
	Age      int       `migu:"size:11"`
	CreateAt time.Time `migu:""`
	UpdateAt time.Time `migu:""`
}

// Scan User all scan
func (e *User) Scan(rows *sql.Rows) error {
	return rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.CreateAt, &e.UpdateAt)
}
