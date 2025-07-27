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

	// Validate request struct
	validationErrors := ValidateStruct(req)
	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": validationErrors,
		})
		return
	}

	// Kiểm tra user có tồn tại không
	if !ValidateUserExists(req.UserID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Kiểm tra tất cả products có tồn tại không
	for _, detail := range req.OrderDetails {
		if !ValidateProductExists(detail.ProductID) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":      "Product not found",
				"product_id": detail.ProductID,
			})
			return
		}
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
