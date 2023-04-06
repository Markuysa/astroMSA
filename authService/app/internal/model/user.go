package model

import "github.com/Markuysa/astroMSA/authService/app/pkg/externalModels"

// User is a structure of user in the system
type User struct {
	ID            uint
	Email         string
	Sign          string
	Name          string
	Password      string
	Notifications bool
	BirthInfo     externalModels.Date
}
