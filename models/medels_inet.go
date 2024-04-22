package models

import (
	"gorm.io/gorm"
)

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}
type BusinessType string

const (
	BusinessTypeRetail      BusinessType = "Retail"
	BusinessTypeWholesale   BusinessType = "Wholesale"
	BusinessTypeManufacture BusinessType = "Manufacture"
	// Add more business types as needed
)

type Register struct {
	Email           string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	Username        string `json:"username" validate:"required,min=3,max=32,viladate-username"`
	Password        string `json:"password" validate:"required,min=3,max=20"`
	LineId          string `json:"lineid" validate:"required,min=3,max=20"`
	TelephoneNumber string `json:"telephonenumber" validate:"required,min=9,max=10"`
	BusinessType    string `json:"businesstype" validate:"required"`
	WebsiteLink     string `json:"website" validate:"required,min=2,max=30,viladate-website"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}
