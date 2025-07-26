package apis

import (
	"gin/config/db"
	models "gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrder tạo order mới cùng các order_detail
func CreateOrder(c *gin.Context) {
	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tạo order
	order := models.Order{
		UserID: req.UserID,
	}
	result := db.GetDB().Create(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Tạo order_detail cho từng sản phẩm
	for _, detail := range req.OrderDetails {
		orderDetail := models.OrderDetail{
			OrderID:   order.ID,
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
		}
		db.GetDB().Create(&orderDetail)
	}

	c.JSON(http.StatusCreated, order)
}
