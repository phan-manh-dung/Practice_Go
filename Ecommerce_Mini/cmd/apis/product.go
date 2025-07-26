package apis

import (
	"gin/config/db"
	models "gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser tạo product mới
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.GetDB().Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// delete product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	result := db.GetDB().Unscoped().Delete(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
	}
}

// update user
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.GetDB().Model(&product).Where("id = ?", id).Updates(product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// get product lấy danh sách tất cả product
func GetProducts(c *gin.Context) {
	var products []models.Product
	result := db.GetDB().Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
