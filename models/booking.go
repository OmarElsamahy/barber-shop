package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UserID   uint
	BarberID uint
	Date     string
	Time     string
	Status   string
}
