package model

import "time"

type User struct {
	Id uint
	Username string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}