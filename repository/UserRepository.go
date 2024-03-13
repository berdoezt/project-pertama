package repository

import (
	"project-pertama/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(newUser model.User) (model.User, error) {
	tx := ur.db.Create(&newUser)
	return newUser, tx.Error
}

func (ur *userRepository) GetByUsername(username string) (model.User, error) {
	var user model.User
	tx := ur.db.First(&user, "username = ?", username)
	return user, tx.Error
}
