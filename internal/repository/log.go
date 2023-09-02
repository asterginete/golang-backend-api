package repository

import (
	"github.com/your_username/go-gin-api/internal/model"
	"github.com/jinzhu/gorm"
)

// Assuming a global DB connection for simplicity. In a real-world scenario, you'd inject the DB connection.
var db *gorm.DB

// AddLog inserts a new log entry into the database.
func AddLog(log model.Log) (model.Log, error) {
	if err := db.Create(&log).Error; err != nil {
		return model.Log{}, err
	}
	return log, nil
}

// GetLogs retrieves all log entries from the database.
func GetLogs() []model.Log {
	var logs []model.Log
	db.Order("created_at desc").Find(&logs)
	return logs
}

// GetLogByID retrieves a single log entry by its ID.
func GetLogByID(id int) (model.Log, error) {
	var log model.Log
	if err := db.First(&log, id).Error; err != nil {
		return model.Log{}, err
	}
	return log, nil
}

// DeleteLog removes a log entry from the database.
func DeleteLog(id int) error {
	var log model.Log
	if err := db.First(&log, id).Error; err != nil {
		return err
	}
	return db.Delete(&log).Error
}

// GetLogsByUserID retrieves all log entries associated with a specific user.
func GetLogsByUserID(userID int) []model.Log {
	var logs []model.Log
	db.Where("user_id = ?", userID).Order("created_at desc").Find(&logs)
	return logs
}

// ... Potentially add other database operations related to logs, like filtering by date range, activity type, etc. ...
