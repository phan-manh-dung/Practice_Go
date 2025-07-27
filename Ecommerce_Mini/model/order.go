package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID           uint          `json:"id" gorm:"primaryKey"`
	UserID       uint          `json:"user_id"` // fk key
	User         User          `json:"user"`    // GORM hiểu đây là liên kết với bảng users
	OrderDetails []OrderDetail `json:"order_details"`
}

// Dùng cho API tạo order
// CreateOrderRequest là struct nhận dữ liệu từ client
// (DTO - Data Transfer Object)
type CreateOrderRequest struct {
	UserID       uint                 `json:"user_id" binding:"required,gt=0" validate:"required,gt=0"`
	OrderDetails []OrderDetailRequest `json:"order_details" binding:"required,min=1" validate:"required,min=1"`
}

type OrderDetailRequest struct {
	ProductID uint `json:"product_id" binding:"required,gt=0" validate:"required,gt=0"`
	Quantity  int  `json:"quantity" binding:"required,gt=0" validate:"required,gt=0"`
}
