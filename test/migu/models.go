package main

import "time"

type User struct {
	ID       int64 `migu:"pk"`
	Name     string
	Email    *string
	Age      int
	CreateAt time.Time
	UpdateAt time.Time
}

type Quest struct {
	ID       int64 `migu:"pk"`
	Title    string
	Detail   string `migu:"size:512"`
	CreateAt time.Time
	UpdateAt time.Time
}
