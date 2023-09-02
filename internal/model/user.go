package model

import "time"

// User represents the structure of a user in our application.
type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"` // Password is hidden in JSON responses
	Email     string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Role      string    `json:"role" gorm:"type:varchar(50);default:user"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// BeforeSave is a GORM hook that can be used to set or modify fields before saving the user.
func (u *User) BeforeSave() error {
	// Any logic you want to run before saving a user
	// For example, you might want to validate the user's data or set default values.
	return nil
}

// AfterSave is a GORM hook that can be used to perform actions after saving the user.
func (u *User) AfterSave() error {
	// Any logic you want to run after saving a user
	// For example, you might want to log the creation of a new user.
	return nil
}

// ... Potentially add other methods or hooks related to the User model ...
