package models

type Order struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint // fk key
	User         User // GORM hiểu đây là liên kết với bảng users
	OrderDetails []OrderDetail
}

// Dùng cho API tạo order
// CreateOrderRequest là struct nhận dữ liệu từ client
// (DTO - Data Transfer Object)
type CreateOrderRequest struct {
	UserID       uint                 `json:"user_id"`
	OrderDetails []OrderDetailRequest `json:"order_details"`
}

type OrderDetailRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
