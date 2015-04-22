package models

import "time"

// User user models
type User struct {
	ID       int64 `migu:"pk"`
	Name     string
	Email    *string
	Age      int
	CreateAt time.Time
	UpdateAt time.Time
}

// Quest quest models
type Quest struct {
	ID       int64 `migu:"pk"`
	Title    string
	Detail   string `migu:"size:512"`
	CreateAt time.Time
	UpdateAt time.Time
}
