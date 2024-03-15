package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UUID      string `gorm:"primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
	Orders    []Order
	DeletedAt gorm.DeletedAt
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.NewString()
	return nil
}
