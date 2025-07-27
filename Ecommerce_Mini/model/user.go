package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string  `json:"name" binding:"required,min=2,max=50" validate:"required,min=2,max=50"`
	SDT    string  `json:"sdt" gorm:"uniqueIndex" binding:"required,len=10" validate:"required,len=10"`
	Orders []Order // 1 user có nhiều order nên dùng slice []Order (1-n) không có ngoặc [] là 1-1
}
