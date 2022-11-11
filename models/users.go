package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	Email     string
	Phone     string
	Password  string
	Langitude float64
	Latitude  float64
}
