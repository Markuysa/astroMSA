package model

import (
	"time"
)

type User struct {
	ID        uint
	Email     string
	Sign      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	BirthInfo time.Time
}
