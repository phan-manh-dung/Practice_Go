package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// slice lưu danh sách người dùng
var users = []User{}
var idCounter = 1

func main() {
	r := gin.Default()

	// create người dùng mới
	r.POST("users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newUser.ID = idCounter
		idCounter++
		users = append(users, newUser)
		c.JSON(http.StatusCreated, newUser)
	})

	// read lấy tất cả người dùng
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// lấy chi tiết người dùng theo id
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			if string(user.ID) == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
	})

	// update user
	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		// log id ra
		fmt.Println("Deleting user with ID:", id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		var updatedUser User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, user := range users {
			if user.ID == idInt {
				users[i].Name = updatedUser.Name
				users[i].Age = updatedUser.Age
				users[i].Email = updatedUser.Email
				c.JSON(http.StatusOK, users[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	// xóa user theo id
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		for i, user := range users {
			if user.ID == idInt {
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	})

	r.Run(":8080")
}
