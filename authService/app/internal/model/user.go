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
	BirthInfo time.Time
}
