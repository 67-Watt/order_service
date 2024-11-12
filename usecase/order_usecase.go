package usecase

import (
	"context"
	"order_service/models"
	"order_service/repository"
)

// OrderUseCase defines the main interface for the Order Use Case
type OrderUseCase interface {
	CreateOrder(ctx context.Context, order models.Order) (uint, error)        // Changed return type to uint
	UpdateOrderStatus(ctx context.Context, orderID uint, status string) error // Updated orderID to uint
	GetOrderDetails(ctx context.Context, orderID uint) (models.Order, error)  // Updated orderID to uint
}

// orderUseCase implements the OrderUseCase interface
type orderUseCase struct {
	repo repository.OrderRepository
}

// NewOrderUseCase creates a new OrderUseCase with the given repository
func NewOrderUseCase(repo repository.OrderRepository) OrderUseCase {
	return &orderUseCase{
		repo: repo,
	}
}

// CreateOrder saves an order and returns the generated OrderID
func (u *orderUseCase) CreateOrder(ctx context.Context, order models.Order) (uint, error) {
	return u.repo.SaveOrder(ctx, order) // Now returning uint directly
}

// UpdateOrderStatus updates the status of an order
func (u *orderUseCase) UpdateOrderStatus(ctx context.Context, orderID uint, status string) error {
	// Add business logic to update order status
	return nil
}

// GetOrderDetails retrieves the details of an order
func (u *orderUseCase) GetOrderDetails(ctx context.Context, orderID uint) (models.Order, error) {
	return u.repo.GetOrder(ctx, orderID)
}
