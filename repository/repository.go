package repository

import (
	"context"
	"gorm.io/gorm"
	"order_service/models"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, order models.Order) (uint, error)  // Updated to return uint
	GetOrder(ctx context.Context, orderID uint) (models.Order, error) // Updated orderID to uint
}

type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new instance of OrderRepository using GORM
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

// SaveOrder inserts a new order record in the database
func (r *orderRepository) SaveOrder(ctx context.Context, order models.Order) (uint, error) {
	if err := r.db.WithContext(ctx).Create(&order).Error; err != nil {
		return 0, err
	}
	return order.OrderID, nil
}

// GetOrder retrieves an order by its ID
func (r *orderRepository) GetOrder(ctx context.Context, orderID uint) (models.Order, error) {
	var order models.Order
	if err := r.db.WithContext(ctx).Preload("OrderDetails").First(&order, orderID).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}
