package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string
	BirthInfo time.Time
	Sign      string
	Name      string
}
