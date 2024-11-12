package models

import (
	"time"
)

type Order struct {
	OrderID       uint          `gorm:"primaryKey;autoIncrement" json:"order_id"` // OrderID now uses uint
	OrderDate     time.Time     `json:"order_date"`
	OrderTime     time.Time     `json:"order_time"`
	EmployeeID    int           `json:"employee_id"`
	CustomerID    int           `json:"customer_id"`
	TableID       int           `json:"table_id"`
	OrderType     string        `json:"order_type"` // Example values: "Dine-In", "Takeout", "Delivery"
	OrderStatusID int           `json:"order_status_id"`
	Discount      float64       `json:"discount"`
	TaxID         int           `json:"tax_id"`
	PaymentMethod string        `json:"payment_method"`
	OrderDetails  []OrderDetail `gorm:"foreignKey:OrderID" json:"order_details"`
}

type OrderDetail struct {
	OrderDetailID       uint    `gorm:"primaryKey;autoIncrement" json:"order_detail_id"`
	OrderID             uint    `json:"order_id"` // OrderID now uses uint
	ItemID              int     `json:"item_id"`
	Quantity            int     `json:"quantity"`
	SpecialInstructions string  `json:"special_instructions"`
	TotalPrice          float64 `json:"total_price"`
}

type Customer struct {
	CustomerID    uint      `gorm:"primaryKey;autoIncrement" json:"customer_id"`
	Name          string    `json:"name"`
	ContactNumber string    `json:"contact_number"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	LoyaltyPoints int       `json:"loyalty_points"`
	LoyaltyLevel  string    `json:"loyalty_level"`
}
