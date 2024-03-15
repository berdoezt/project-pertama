package repository

import (
	"project-pertama/model"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) GetAll() ([]model.Order, error) {
	var orders = []model.Order{}

	tx := or.db.Find(&orders)
	return orders, tx.Error
}

func (or *orderRepository) GetAllByUserId(userId string) ([]model.Order, error) {
	var orders = []model.Order{}

	tx := or.db.Find(&orders, "user_uuid = ?", userId)
	return orders, tx.Error
}

func (or *orderRepository) Create(newOrder model.Order) (model.Order, error) {
	tx := or.db.Create(&newOrder)
	return newOrder, tx.Error
}

func (or *orderRepository) Delete(uuid string) error {
	tx := or.db.Delete(&model.Order{}, "uuid = ?", uuid)
	return tx.Error
}
