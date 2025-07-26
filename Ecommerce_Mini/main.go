package main

import (
	"gin/cmd/apis"
	"gin/config/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Kết nối database
	db.ConnectDatabase()

	// Khởi tạo Gin router
	r := gin.Default()

	// Lấy instance database để sử dụng
	database := db.GetDB()
	_ = database // Sử dụng database để tránh warning rỗng k dùng đến

	r.POST("/users", apis.CreateUser)
	r.GET("/users", apis.GetUsers)
	r.DELETE("/users/:id", apis.DeleteUser)
	r.PUT("/users/:id", apis.UpdateUser)

	// products
	r.POST("/products", apis.CreateProduct)
	r.GET("/products", apis.GetProducts)
	r.DELETE("/products/:id", apis.DeleteProduct)
	r.PUT("/products/:id", apis.UpdateProduct)

	// order
	r.POST("/orders", apis.CreateOrder)

	r.Run(":8080")
}
