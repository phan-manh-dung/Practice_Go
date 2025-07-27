package models

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" binding:"required,min=2,max=100" validate:"required,min=2,max=100"`
	Description string  `json:"description" binding:"required,min=10,max=500" validate:"required,min=10,max=500"`
	Price       float64 `json:"price" binding:"required,gt=0" validate:"required,gt=0"`
}
