package models

type OrderDetail struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	Order     Order
	ProductID uint
	Product   Product
	Quantity  int
}
