package models

import (
	"gorm.io/gorm"
	"Server/src/config"
	"fmt"
)

var db *gorm.DB
type SignUp struct{
	gorm.Model
	
	FirstName string 	`gorm:"not null;" json:"firstname"`
	LastName  string 	`gorm:"not null;" json:"lastname"`
	Email     string	`gorm:"not null;" json:"email"`
	Password  string		`gorm:"not null;" json:"password"`
	ConPassword  string	`gorm:"not null;" json:"conpassword"`
}


func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&SignUp{})
	fmt.Println("Successfully connected!")
}

func (s *SignUp) StoreSignupDetails() *SignUp{
	// db.NewRecord(s)
	db.Create(&s)
	return s
}