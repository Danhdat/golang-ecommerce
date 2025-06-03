package controller

import (
	"net/http"
	config "store1/data"
	models "store1/model"

	"github.com/gin-gonic/gin"
)

// Tạo category
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&category)
	c.JSON(http.StatusOK, category)
}

// Lấy tất cả category
func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := config.DB.Find(&categories).Error
	return categories, err
}

// Lấy 1 category theo ID
func GetCategoryByID(id string) ([]models.Category, error) {
	var category []models.Category
	err := config.DB.First(&category, id).Error
	return category, err
}

// Cập nhật category
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy danh mục"})
		return
	}

	var update models.Category
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&category).Updates(update)
	c.JSON(http.StatusOK, category)
}

// Xoá category
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy danh mục"})
		return
	}

	config.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"message": "Đã xoá danh mục"})
}
