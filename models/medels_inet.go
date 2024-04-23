package models

import (
	// "time"
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

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}
type ResultData struct {
	Data        []DogsRes `json:"data"`
	Name        string    `json:"name"`
	Count       int       `json:"count"`
	Sum_Red     int       `json:"sum_red"`
	Sum_Green   int       `json:"sum_green"`
	Sum_Pink    int       `json:"sum_pink"`
	Sum_NoColor int       `json:"sum_nocolor"`
}

// 7.0.1 Create Database for Store Company Data
type Company struct {
	gorm.Model
	CompanyName    string `json:"company_name" validate:"required,min=3,max=20"`
	CompanyAddress string `json:"company_address" validate:"required,min=9,max=150"`
	Tel            string `json:"tel" validate:"required,min=9,max=10"`
	Email          string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	WebsiteLink    string `json:"website" validate:"required,min=2,max=30,website"`
	Employee       int    `json:"emp_amount"`
}

type Profile struct {
	gorm.Model
	EmployeeID string `json:"emp_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Birthday   string `json:"birthday" validate:"required"`
	Age        int    `json:"age"  validate:"required"`
	Email      string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	Tel        string `json:"tel" validate:"required,min=9,max=10"`
}

type ProfileRes struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Type string `json:"type"`
}
type ResultProfileData struct {
	Data         []ProfileRes `json:"data"`
	Name         string       `json:"name"`
	Count        int          `json:"count"`
	GenZ         int          `json:"gen_z"`
	GenX         int          `json:"gen_x"`
	GenY         int          `json:"gen_y"`
	BabyBoomer   int          `json:"baby_boomer"`
	GIGeneration int          `json:"gi_generation"`
}
