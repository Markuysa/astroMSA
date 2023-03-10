package model

import (
	"time"
)

type User struct {
	ID        uint
	Email     string
	Sign      string
	Name      string
	Password  string
	CreatedAt time.Time
	BirthInfo time.Time
}
