package models

import "gorm.io/gorm"

type Barber struct {
	gorm.Model
	Name     string
	Location string
	Phone    string
}
