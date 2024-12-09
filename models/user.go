package models

import "gorm.io/gorm"

// User model (Make sure the correct column is used here)
type User struct {
	gorm.Model
	UserID   uint   `json:"user_id"` // Ensure this matches your table schema
	UserName string `json:"user_name"`
}

// TableName function to specify the correct table name
func (User) TableName() string {
	return "user_s" // Make sure the correct table name is used
}
