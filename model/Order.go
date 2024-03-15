package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	UUID        string `gorm:"primaryKey"`
	UserUUID    string
	TotalAmount int `json:"total_amount"`
	DeletedAt   gorm.DeletedAt
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.UUID = uuid.NewString()
	return nil
}

type OrderRequest struct {
	TotalAmount int `json:"total_amount"`
}

type OrderResponse struct {
	UUID        string `json:"order_id"`
	TotalAmount int    `json:"total_amount"`
}
