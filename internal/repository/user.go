package repository

import (
	"github.com/your_username/go-gin-api/internal/model"
	"github.com/jinzhu/gorm"
)

// Assuming a global DB connection for simplicity. In a real-world scenario, you'd inject the DB connection.
var db *gorm.DB

// AddUser inserts a new user into the database.
func AddUser(user model.User) (model.User, error) {
	if err := db.Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GetUserByUsername retrieves a user by their username.
func GetUserByUsername(username string) (model.User, error) {
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GetUserByID retrieves a single user by their ID.
func GetUserByID(id int) (model.User, error) {
	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

// UpdateUser updates a user's details in the database.
func UpdateUser(id int, updatedUser model.User) (model.User, error) {
	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return model.User{}, err
	}

	// Update fields
	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	user.Role = updatedUser.Role

	if updatedUser.Password != "" {
		user.Password = updatedUser.Password
	}

	if err := db.Save(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

// DeleteUser removes a user from the database.
func DeleteUser(id int) error {
	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return err
	}
	return db.Delete(&user).Error
}

// GetUsers retrieves all users from the database.
func GetUsers() []model.User {
	var users []model.User
	db.Find(&users)
	return users
}

// ... Potentially add other database operations related to users, like filtering, searching, etc. ...
