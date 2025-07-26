package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string
	SDT    string  `gorm:"uniqueIndex"`
	Orders []Order // 1 user có nhiều order nên dùng slice []Order (1-n) không có ngoặc [] là 1-1
}
