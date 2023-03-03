package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email     string
	Sign      string
	Name      string
	BirthInfo time.Time
}
