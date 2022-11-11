package models

import "gorm.io/gorm"

type Pharmacy struct {
	gorm.Model
	Name      string
	Location  string
	Phone     string
	Email     string
	Langitude float64
	Latitude  float64
}

type Drug struct {
	gorm.Model
	Name        string
	Price       float64
	Quantity    int
	PharmacyID  uint
	Description string
}
