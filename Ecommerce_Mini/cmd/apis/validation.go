package apis

import (
	"gin/config/db"
	models "gin/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidationErrorResponse trả về lỗi validation dạng JSON
type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

// ValidateRequest middleware để validate request
func ValidateRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// ValidateStruct validate một struct và trả về lỗi chi tiết
func ValidateStruct(s interface{}) []ValidationErrorResponse {
	var errors []ValidationErrorResponse

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationErrorResponse
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.Message = getErrorMessage(err.Tag(), err.Field())
			errors = append(errors, element)
		}
	}

	return errors
}

// getErrorMessage trả về message lỗi tùy chỉnh
func getErrorMessage(tag string, field string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "min":
		return field + " must be at least " + tag + " characters"
	case "max":
		return field + " must be at most " + tag + " characters"
	case "len":
		return field + " must be exactly " + tag + " characters"
	case "gt":
		return field + " must be greater than " + tag
	case "gte":
		return field + " must be greater than or equal to " + tag
	case "lt":
		return field + " must be less than " + tag
	case "lte":
		return field + " must be less than or equal to " + tag
	default:
		return field + " is invalid"
	}
}

// ValidateID validate ID parameter
func ValidateID(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// ValidateUserExists kiểm tra user có tồn tại không
func ValidateUserExists(userID uint) bool {
	var count int64
	db.GetDB().Model(&models.User{}).Where("id = ?", userID).Count(&count)
	return count > 0
}

// ValidateProductExists kiểm tra product có tồn tại không
func ValidateProductExists(productID uint) bool {
	var count int64
	db.GetDB().Model(&models.Product{}).Where("id = ?", productID).Count(&count)
	return count > 0
}
