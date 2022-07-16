package models

import (
	"gorm.io/gorm"
)

//model user
type User struct {
	gorm.Model
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	UserName  string `gorm:"unique;not null"`
	Gender string `json:"gender"`
	Rol  string `json:"roles"`
	Password  string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	
}