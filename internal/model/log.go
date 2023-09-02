package model

import "time"

// Log represents the structure of a log entry in our application.
type Log struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"user_id" gorm:"type:int;not null"`
	Activity  string    `json:"activity" gorm:"type:text;not null"`
	IPAddress string    `json:"ip_address" gorm:"type:varchar(50)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// BeforeSave is a GORM hook that can be used to set or modify fields before saving the log entry.
func (l *Log) BeforeSave() error {
	// Any logic you want to run before saving a log entry
	// For example, you might want to validate the log's data or set default values.
	return nil
}

// AfterSave is a GORM hook that can be used to perform actions after saving the log entry.
func (l *Log) AfterSave() error {
	// Any logic you want to run after saving a log entry
	// For example, you might want to notify administrators of certain activities.
	return nil
}

// ... Potentially add other methods or hooks related to the Log model ...
