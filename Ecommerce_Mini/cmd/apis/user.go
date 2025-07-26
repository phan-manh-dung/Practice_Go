package apis

import (
	"gin/config/db"
	models "gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser tạo user mới
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.GetDB().Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// delete user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := db.GetDB().Unscoped().Delete(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}

// update user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.GetDB().Model(&user).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUsers lấy danh sách tất cả users
func GetUsers(c *gin.Context) {
	var users []models.User
	result := db.GetDB().Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
