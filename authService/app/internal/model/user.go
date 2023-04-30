package model

import (
	"gorm.io/gorm"
	"time"
)

// User is a structure of user in the system
type User struct {
	gorm.Model
	ID            uint
	Email         string
	Sign          string
	Name          string
	Password      string
	Notifications bool
	BirthInfo     time.Time
}
